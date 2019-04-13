package main

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/rekby/zapcontext"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme"
	"sync"
	"time"
)

const (
	tlsAlpn01 = "tls-alpn-01"
)

var (
	allowedChallenges = []string{tlsAlpn01}
)

var ErrCacheMiss = errors.New("lets proxy: certificate cache miss")
var haveNoCert = errors.New("have no certificate for domain")
var notImplementedError = errors.New("not implemented yet")

type Cache interface {
	// Get returns a certificate data for the specified key.
	// If there's no such key, Get returns ErrCacheMiss.
	Get(ctx context.Context, key string) ([]byte, error)

	// Put stores the data in the cache under the specified key.
	// Underlying implementations may use any data storage format,
	// as long as the reverse operation, Get, results in the original data.
	Put(ctx context.Context, key string, data []byte) error

	// Delete removes a certificate data from the cache under the specified key.
	// If there's no such key in the cache, Delete returns nil.
	Delete(ctx context.Context, key string) error
}

// Interface inspired to https://godoc.org/golang.org/x/crypto/acme/autocert#Manager but not compatible gurantee
type Manager struct {
	GetCertContext          context.Context // base context for use in GetCertificate - use for logging and cancel.
	CertificateIssueTimeout time.Duration

	// Client is used to perform low-level operations, such as account registration
	// and requesting new certificates.
	//
	// If Client is nil, a zero-value acme.Client is used with acme.LetsEncryptURL
	// as directory endpoint. If the Client.Key is nil, a new ECDSA P-256 key is
	// generated and, if Cache is not nil, stored in cache.
	//
	// Mutating the field after the first call of GetCertificate method will have no effect.
	Client *acme.Client

	tokensMu sync.RWMutex
	tokens   map[string]*tls.Certificate
}

// GetCertificate implements the tls.Config.GetCertificate hook.
func (m *Manager) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	// TODO: get context of connection

	domain := hello.ServerName
	logger := zc.L(m.GetCertContext).With(logDomain(domain))
	logger.Info("Get certificate")
	if isTlsAlpn01Hello(hello) {
		logger.Debug("It is token request domain.")
		m.tokensMu.RLock()
		cert := m.tokens[domain]
		m.tokensMu.RUnlock()

		if cert == nil {
			logger.Warn("Doesn't have token for request domain")
			return nil, haveNoCert
		}
	}

	// TODO: check cache

	// TODO: check domain
	certIssueContext, cancelFunc := context.WithTimeout(m.GetCertContext, m.CertificateIssueTimeout)
	defer cancelFunc()

	// TODO: receive cert for domain and subdomains same time
	res, err := m.createCertificateForDomain(certIssueContext, domain)
	if err == nil {
		logger.Info("Certificate issued.", zap.Strings("cert_domains", res.Leaf.DNSNames),
			zap.Time("expire", res.Leaf.NotAfter))
		return res, nil
	} else {
		logger.Warn("Can't issue certificate", zap.Error(err))
		return nil, haveNoCert
	}
}

func (m *Manager) createCertificateForDomain(ctx context.Context, domain string) (*tls.Certificate, error) {
	logger := zc.L(ctx).With(logDomain(domain))
	ctx = zc.WithLogger(ctx, logger)

	// TODO: syncronize many goroutines get certificate for one domain same time
	err := m.authorizeDomain(ctx, domain)
	if err == nil {
		logger.Debug("Domain authorized.")
	} else {
		logger.Warn("Domain doesn't authorized.", zap.Error(err))
		return nil, err
	}

	res, err := m.issueCertificate(ctx, domain)
	if err == nil {
		logger.Debug("Certificate created.")
	} else {
		logger.Warn("Can't issue certificate", zap.Error(err))
	}
	return res, err
}

func (m *Manager) authorizeDomain(ctx context.Context, domain string) error {
	logger := zc.L(ctx).With(logDomain(domain))
	pendingAuthorizations := make(map[string]struct{})
	defer func() {
		if len(pendingAuthorizations) == 0 {
			logger.Debug("No pending authorizations to clean")
			return
		}

		var authUries []string
		for k := range pendingAuthorizations {
			authUries = append(authUries, k)
		}

		logger.Info("Start detached process for revoke pending authorizations", zap.Strings("uries", authUries))

		revokeContext := context.Background()
		revokeContext = zc.WithLogger(revokeContext, logger.With(zap.Bool("detached", true)))
		go m.revokePendingAuthorizations(revokeContext, authUries)
	}()

	var nextChallengeTypeIndex int
	for {
		// Start domain authorization and get the challenge.
		authz, err := m.Client.Authorize(ctx, domain)
		if err == nil {
			logger.Debug("Got authorization description.", zap.Reflect("auth_object", authz))
		} else {
			logger.Error("Can't get domain authorization description", zap.Error(err))
			return err
		}

		// No point in accepting challenges if the authorization status
		// is in a final state.
		switch authz.Status {
		case acme.StatusValid:
			logger.Debug("Domain already authorized")
			return nil // already authorized
		case acme.StatusInvalid:
			logger.Warn("Domain has invalid authorization", zap.String("auth_uri", authz.URI))
			return errors.New("invalid authorization status")
		}

		pendingAuthorizations[authz.URI] = struct{}{}

		// Pick the next preferred challenge.
		var chal *acme.Challenge
		for chal == nil && nextChallengeTypeIndex < len(allowedChallenges) {
			logger.Debug("Check if accept challenge",
				zap.String("challenge", allowedChallenges[nextChallengeTypeIndex]))
			chal = pickChallenge(allowedChallenges[nextChallengeTypeIndex], authz.Challenges)
			nextChallengeTypeIndex++
		}
		if chal == nil {
			logger.Warn("Unable to authorize domain. No compatible challenges.")
			return errors.New("unable authorize domain")
		} else {
			logger.Debug("Select challenge for authorize", zap.Reflect("challenge", chal))
		}

		cleanup, err := m.fulfill(ctx, chal, domain)
		if err != nil {
			logger.Error("Can't set domain token", zap.Reflect("chal", chal), zap.Error(err))
			continue
		}
		//noinspection GoDeferInLoop
		defer cleanup(ctx)

		receivedChallenge, err := m.Client.Accept(ctx, chal)
		if err == nil {
			logger.Debug("Receive authorize answer", zap.Reflect("challenge", receivedChallenge))
		} else {
			logger.Error("Can't authorize domain", zap.Reflect("challenge", receivedChallenge), zap.Error(err))
			continue
		}

		receivedAuth, err := m.Client.WaitAuthorization(ctx, authz.URI)
		if err == nil {
			logger.Debug("Receive domain authorization", zap.Reflect("authorization", receivedAuth))
		} else {
			logger.Error("Dont receive wait authorization", zap.Reflect("authorizarion", receivedAuth), zap.Error(err))
			continue
		}

		logger.Info("Domain authorized")
		delete(pendingAuthorizations, authz.URI)
		return nil
	}
}

func (m *Manager) issueCertificate(ctx context.Context, domain string) (*tls.Certificate, error) {
	// TODO: issue certificate for many domains same time
	return nil, notImplementedError
}

func (m *Manager) revokePendingAuthorizations(revokeContext context.Context, strings []string) {
	// TODO:
}

func (m *Manager) fulfill(ctx context.Context, challenge *acme.Challenge, domain string) (func(context.Context), error) {
	logger := zc.L(ctx)
	switch challenge.Type {
	case tlsAlpn01:
		cert, err := m.Client.TLSALPN01ChallengeCert(challenge.Token, domain)
		if err != nil {
			return nil, err
		}
		m.putCertToken(ctx, domain, &cert)
		return func(ctx context.Context) { go m.deleteCertToken(ctx, domain) }, nil
	default:
		logger.Error("Unknow challenge type", zap.Reflect("challenge", challenge))
		return nil, errors.New("unknown challenge type")
	}
}

func (m *Manager) putCertToken(ctx context.Context, key string, certificate *tls.Certificate) {
	logger := zc.L(ctx)
	logger.Debug("Put cert token", zap.String("key", key))

	var overwriteToken bool

	defer func() {
		if overwriteToken {
			logger.Warn("Unexpected cert token already exist", zap.String("key", key))
		}
	}()

	m.tokensMu.Lock()
	defer m.tokensMu.Unlock()

	_, overwriteToken = m.tokens[key]
	m.tokens[key] = certificate
}

func (m *Manager) deleteCertToken(ctx context.Context, key string) {
	logger := zc.L(ctx)
	logger.Debug("Delete cert token", zap.String("key", key))

	var exist bool
	defer func() {
		if !exist {
			logger.Warn("Cert token for delete doesn't exist", zap.String("key", key))
		}
	}()

	m.tokensMu.Lock()
	defer m.tokensMu.Unlock()

	_, exist = m.tokens[key]
	delete(m.tokens, key)
}

func isTlsAlpn01Hello(hello *tls.ClientHelloInfo) bool {
	return len(hello.SupportedProtos) == 1 && hello.SupportedProtos[0] == acme.ALPNProto
}

func pickChallenge(typ string, chal []*acme.Challenge) *acme.Challenge {
	for _, c := range chal {
		if c.Type == typ {
			return c
		}
	}
	return nil
}
