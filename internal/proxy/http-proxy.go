package proxy

import (
	"context"
	"errors"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rekby/lets-proxy2/internal/log"

	"go.uber.org/zap"

	zc "github.com/rekby/zapcontext"
)

type HTTPProxy struct {
	GetDestination       func(ctx context.Context, remoteAddr string) (addr string, err error)
	GetContext           func(req *http.Request) (context.Context, error)
	HandleHTTPValidation func(w http.ResponseWriter, r *http.Request) bool

	ctx              context.Context
	listener         net.Listener
	httpReverseProxy httputil.ReverseProxy
}

func NewHTTPProxy(ctx context.Context, listener net.Listener) *HTTPProxy {
	res := &HTTPProxy{
		GetDestination: getDestination,
		HandleHTTPValidation: func(_ http.ResponseWriter, _ *http.Request) bool {
			return false
		},
		GetContext: getContext,
		listener:   listener,
		ctx:        ctx,
	}
	res.httpReverseProxy.Director = res.director

	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if !res.HandleHTTPValidation(writer, request) {
			res.httpReverseProxy.ServeHTTP(writer, request)
		}
	})
	httpServer := http.Server{}
	httpServer.Handler = mux

	go func() {
		<-ctx.Done()
		err := httpServer.Close()
		log.DebugErrorCtx(ctx, err, "Http builtin reverse proxy stop because context cancelled")
	}()

	go func() {
		zc.L(ctx).Info("Http builtin reverse proxy start")
		err := httpServer.Serve(res.listener)
		if err == http.ErrServerClosed {
			err = nil
		}
		log.DebugErrorCtx(ctx, err, "Http builtin reverse proxy stop")
	}()

	return res
}

func (p *HTTPProxy) SetTransport(transport http.RoundTripper) {
	p.httpReverseProxy.Transport = transport
}

func getDestination(_ context.Context, remoteAddr string) (addr string, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		return "", errors.New("default get destination accept only tcp addresses")
	}
	tcpAddr.Port = 80
	return tcpAddr.String(), nil
}

func getContext(_ *http.Request) (context.Context, error) {
	return zc.WithLogger(context.Background(), zap.NewNop()), nil
}

func (p *HTTPProxy) director(request *http.Request) {
	ctx := p.getContext(request)
	if request.URL == nil {
		request.URL = &url.URL{}
	}
	dest, err := p.GetDestination(ctx, request.RemoteAddr)
	log.DebugErrorCtx(ctx, err, "Get destination", zap.String("dest_addr", dest))
	request.URL.Scheme = "http"
	request.URL.Host = dest // If err != nil and dest invalid - is ok, because it will error proxy
}

func (p *HTTPProxy) getContext(req *http.Request) context.Context {
	ctx, err := p.GetContext(req)
	if err == nil {
		return ctx
	}

	connectionID := rand.Int63()
	logger := zc.L(p.ctx).With(zap.Int64("connection_id", connectionID))
	logger.DPanic("Http proxy can't receive proxy context. Create own connection_id.")
	ctx = zc.WithLogger(p.ctx, logger)
	return ctx
}
