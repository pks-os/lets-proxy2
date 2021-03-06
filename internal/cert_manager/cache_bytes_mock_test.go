package cert_manager

// Code generated by http://github.com/gojuno/minimock (3.0.6). DO NOT EDIT.

//go:generate minimock -i github.com/rekby/lets-proxy2/internal/cache.Bytes -o ./cache_bytes_mock_test.go

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// BytesMock implements cache.Bytes
type BytesMock struct {
	t minimock.Tester

	funcDelete          func(ctx context.Context, key string) (err error)
	inspectFuncDelete   func(ctx context.Context, key string)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mBytesMockDelete

	funcGet          func(ctx context.Context, key string) (ba1 []byte, err error)
	inspectFuncGet   func(ctx context.Context, key string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mBytesMockGet

	funcPut          func(ctx context.Context, key string, data []byte) (err error)
	inspectFuncPut   func(ctx context.Context, key string, data []byte)
	afterPutCounter  uint64
	beforePutCounter uint64
	PutMock          mBytesMockPut
}

// NewBytesMock returns a mock for cache.Bytes
func NewBytesMock(t minimock.Tester) *BytesMock {
	m := &BytesMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mBytesMockDelete{mock: m}
	m.DeleteMock.callArgs = []*BytesMockDeleteParams{}

	m.GetMock = mBytesMockGet{mock: m}
	m.GetMock.callArgs = []*BytesMockGetParams{}

	m.PutMock = mBytesMockPut{mock: m}
	m.PutMock.callArgs = []*BytesMockPutParams{}

	return m
}

type mBytesMockDelete struct {
	mock               *BytesMock
	defaultExpectation *BytesMockDeleteExpectation
	expectations       []*BytesMockDeleteExpectation

	callArgs []*BytesMockDeleteParams
	mutex    sync.RWMutex
}

// BytesMockDeleteExpectation specifies expectation struct of the Bytes.Delete
type BytesMockDeleteExpectation struct {
	mock    *BytesMock
	params  *BytesMockDeleteParams
	results *BytesMockDeleteResults
	Counter uint64
}

// BytesMockDeleteParams contains parameters of the Bytes.Delete
type BytesMockDeleteParams struct {
	ctx context.Context
	key string
}

// BytesMockDeleteResults contains results of the Bytes.Delete
type BytesMockDeleteResults struct {
	err error
}

// Expect sets up expected params for Bytes.Delete
func (mmDelete *mBytesMockDelete) Expect(ctx context.Context, key string) *mBytesMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BytesMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &BytesMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &BytesMockDeleteParams{ctx, key}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Bytes.Delete
func (mmDelete *mBytesMockDelete) Inspect(f func(ctx context.Context, key string)) *mBytesMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for BytesMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Bytes.Delete
func (mmDelete *mBytesMockDelete) Return(err error) *BytesMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BytesMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &BytesMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &BytesMockDeleteResults{err}
	return mmDelete.mock
}

//Set uses given function f to mock the Bytes.Delete method
func (mmDelete *mBytesMockDelete) Set(f func(ctx context.Context, key string) (err error)) *BytesMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Bytes.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Bytes.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Bytes.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mBytesMockDelete) When(ctx context.Context, key string) *BytesMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BytesMock.Delete mock is already set by Set")
	}

	expectation := &BytesMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &BytesMockDeleteParams{ctx, key},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Bytes.Delete return parameters for the expectation previously defined by the When method
func (e *BytesMockDeleteExpectation) Then(err error) *BytesMock {
	e.results = &BytesMockDeleteResults{err}
	return e.mock
}

// Delete implements cache.Bytes
func (mmDelete *BytesMock) Delete(ctx context.Context, key string) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, key)
	}

	mm_params := &BytesMockDeleteParams{ctx, key}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := BytesMockDeleteParams{ctx, key}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("BytesMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the BytesMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, key)
	}
	mmDelete.t.Fatalf("Unexpected call to BytesMock.Delete. %v %v", ctx, key)
	return
}

// DeleteAfterCounter returns a count of finished BytesMock.Delete invocations
func (mmDelete *BytesMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of BytesMock.Delete invocations
func (mmDelete *BytesMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to BytesMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mBytesMockDelete) Calls() []*BytesMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*BytesMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *BytesMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *BytesMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BytesMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BytesMock.Delete")
		} else {
			m.t.Errorf("Expected call to BytesMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to BytesMock.Delete")
	}
}

type mBytesMockGet struct {
	mock               *BytesMock
	defaultExpectation *BytesMockGetExpectation
	expectations       []*BytesMockGetExpectation

	callArgs []*BytesMockGetParams
	mutex    sync.RWMutex
}

// BytesMockGetExpectation specifies expectation struct of the Bytes.Get
type BytesMockGetExpectation struct {
	mock    *BytesMock
	params  *BytesMockGetParams
	results *BytesMockGetResults
	Counter uint64
}

// BytesMockGetParams contains parameters of the Bytes.Get
type BytesMockGetParams struct {
	ctx context.Context
	key string
}

// BytesMockGetResults contains results of the Bytes.Get
type BytesMockGetResults struct {
	ba1 []byte
	err error
}

// Expect sets up expected params for Bytes.Get
func (mmGet *mBytesMockGet) Expect(ctx context.Context, key string) *mBytesMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BytesMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BytesMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &BytesMockGetParams{ctx, key}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Bytes.Get
func (mmGet *mBytesMockGet) Inspect(f func(ctx context.Context, key string)) *mBytesMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for BytesMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Bytes.Get
func (mmGet *mBytesMockGet) Return(ba1 []byte, err error) *BytesMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BytesMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BytesMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &BytesMockGetResults{ba1, err}
	return mmGet.mock
}

//Set uses given function f to mock the Bytes.Get method
func (mmGet *mBytesMockGet) Set(f func(ctx context.Context, key string) (ba1 []byte, err error)) *BytesMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Bytes.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Bytes.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Bytes.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mBytesMockGet) When(ctx context.Context, key string) *BytesMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BytesMock.Get mock is already set by Set")
	}

	expectation := &BytesMockGetExpectation{
		mock:   mmGet.mock,
		params: &BytesMockGetParams{ctx, key},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Bytes.Get return parameters for the expectation previously defined by the When method
func (e *BytesMockGetExpectation) Then(ba1 []byte, err error) *BytesMock {
	e.results = &BytesMockGetResults{ba1, err}
	return e.mock
}

// Get implements cache.Bytes
func (mmGet *BytesMock) Get(ctx context.Context, key string) (ba1 []byte, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, key)
	}

	mm_params := &BytesMockGetParams{ctx, key}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ba1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := BytesMockGetParams{ctx, key}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("BytesMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the BytesMock.Get")
		}
		return (*mm_results).ba1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, key)
	}
	mmGet.t.Fatalf("Unexpected call to BytesMock.Get. %v %v", ctx, key)
	return
}

// GetAfterCounter returns a count of finished BytesMock.Get invocations
func (mmGet *BytesMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of BytesMock.Get invocations
func (mmGet *BytesMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to BytesMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mBytesMockGet) Calls() []*BytesMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*BytesMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *BytesMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *BytesMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BytesMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BytesMock.Get")
		} else {
			m.t.Errorf("Expected call to BytesMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to BytesMock.Get")
	}
}

type mBytesMockPut struct {
	mock               *BytesMock
	defaultExpectation *BytesMockPutExpectation
	expectations       []*BytesMockPutExpectation

	callArgs []*BytesMockPutParams
	mutex    sync.RWMutex
}

// BytesMockPutExpectation specifies expectation struct of the Bytes.Put
type BytesMockPutExpectation struct {
	mock    *BytesMock
	params  *BytesMockPutParams
	results *BytesMockPutResults
	Counter uint64
}

// BytesMockPutParams contains parameters of the Bytes.Put
type BytesMockPutParams struct {
	ctx  context.Context
	key  string
	data []byte
}

// BytesMockPutResults contains results of the Bytes.Put
type BytesMockPutResults struct {
	err error
}

// Expect sets up expected params for Bytes.Put
func (mmPut *mBytesMockPut) Expect(ctx context.Context, key string, data []byte) *mBytesMockPut {
	if mmPut.mock.funcPut != nil {
		mmPut.mock.t.Fatalf("BytesMock.Put mock is already set by Set")
	}

	if mmPut.defaultExpectation == nil {
		mmPut.defaultExpectation = &BytesMockPutExpectation{}
	}

	mmPut.defaultExpectation.params = &BytesMockPutParams{ctx, key, data}
	for _, e := range mmPut.expectations {
		if minimock.Equal(e.params, mmPut.defaultExpectation.params) {
			mmPut.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPut.defaultExpectation.params)
		}
	}

	return mmPut
}

// Inspect accepts an inspector function that has same arguments as the Bytes.Put
func (mmPut *mBytesMockPut) Inspect(f func(ctx context.Context, key string, data []byte)) *mBytesMockPut {
	if mmPut.mock.inspectFuncPut != nil {
		mmPut.mock.t.Fatalf("Inspect function is already set for BytesMock.Put")
	}

	mmPut.mock.inspectFuncPut = f

	return mmPut
}

// Return sets up results that will be returned by Bytes.Put
func (mmPut *mBytesMockPut) Return(err error) *BytesMock {
	if mmPut.mock.funcPut != nil {
		mmPut.mock.t.Fatalf("BytesMock.Put mock is already set by Set")
	}

	if mmPut.defaultExpectation == nil {
		mmPut.defaultExpectation = &BytesMockPutExpectation{mock: mmPut.mock}
	}
	mmPut.defaultExpectation.results = &BytesMockPutResults{err}
	return mmPut.mock
}

//Set uses given function f to mock the Bytes.Put method
func (mmPut *mBytesMockPut) Set(f func(ctx context.Context, key string, data []byte) (err error)) *BytesMock {
	if mmPut.defaultExpectation != nil {
		mmPut.mock.t.Fatalf("Default expectation is already set for the Bytes.Put method")
	}

	if len(mmPut.expectations) > 0 {
		mmPut.mock.t.Fatalf("Some expectations are already set for the Bytes.Put method")
	}

	mmPut.mock.funcPut = f
	return mmPut.mock
}

// When sets expectation for the Bytes.Put which will trigger the result defined by the following
// Then helper
func (mmPut *mBytesMockPut) When(ctx context.Context, key string, data []byte) *BytesMockPutExpectation {
	if mmPut.mock.funcPut != nil {
		mmPut.mock.t.Fatalf("BytesMock.Put mock is already set by Set")
	}

	expectation := &BytesMockPutExpectation{
		mock:   mmPut.mock,
		params: &BytesMockPutParams{ctx, key, data},
	}
	mmPut.expectations = append(mmPut.expectations, expectation)
	return expectation
}

// Then sets up Bytes.Put return parameters for the expectation previously defined by the When method
func (e *BytesMockPutExpectation) Then(err error) *BytesMock {
	e.results = &BytesMockPutResults{err}
	return e.mock
}

// Put implements cache.Bytes
func (mmPut *BytesMock) Put(ctx context.Context, key string, data []byte) (err error) {
	mm_atomic.AddUint64(&mmPut.beforePutCounter, 1)
	defer mm_atomic.AddUint64(&mmPut.afterPutCounter, 1)

	if mmPut.inspectFuncPut != nil {
		mmPut.inspectFuncPut(ctx, key, data)
	}

	mm_params := &BytesMockPutParams{ctx, key, data}

	// Record call args
	mmPut.PutMock.mutex.Lock()
	mmPut.PutMock.callArgs = append(mmPut.PutMock.callArgs, mm_params)
	mmPut.PutMock.mutex.Unlock()

	for _, e := range mmPut.PutMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPut.PutMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPut.PutMock.defaultExpectation.Counter, 1)
		mm_want := mmPut.PutMock.defaultExpectation.params
		mm_got := BytesMockPutParams{ctx, key, data}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPut.t.Errorf("BytesMock.Put got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPut.PutMock.defaultExpectation.results
		if mm_results == nil {
			mmPut.t.Fatal("No results are set for the BytesMock.Put")
		}
		return (*mm_results).err
	}
	if mmPut.funcPut != nil {
		return mmPut.funcPut(ctx, key, data)
	}
	mmPut.t.Fatalf("Unexpected call to BytesMock.Put. %v %v %v", ctx, key, data)
	return
}

// PutAfterCounter returns a count of finished BytesMock.Put invocations
func (mmPut *BytesMock) PutAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPut.afterPutCounter)
}

// PutBeforeCounter returns a count of BytesMock.Put invocations
func (mmPut *BytesMock) PutBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPut.beforePutCounter)
}

// Calls returns a list of arguments used in each call to BytesMock.Put.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPut *mBytesMockPut) Calls() []*BytesMockPutParams {
	mmPut.mutex.RLock()

	argCopy := make([]*BytesMockPutParams, len(mmPut.callArgs))
	copy(argCopy, mmPut.callArgs)

	mmPut.mutex.RUnlock()

	return argCopy
}

// MinimockPutDone returns true if the count of the Put invocations corresponds
// the number of defined expectations
func (m *BytesMock) MinimockPutDone() bool {
	for _, e := range m.PutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPut != nil && mm_atomic.LoadUint64(&m.afterPutCounter) < 1 {
		return false
	}
	return true
}

// MinimockPutInspect logs each unmet expectation
func (m *BytesMock) MinimockPutInspect() {
	for _, e := range m.PutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BytesMock.Put with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutCounter) < 1 {
		if m.PutMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BytesMock.Put")
		} else {
			m.t.Errorf("Expected call to BytesMock.Put with params: %#v", *m.PutMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPut != nil && mm_atomic.LoadUint64(&m.afterPutCounter) < 1 {
		m.t.Error("Expected call to BytesMock.Put")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BytesMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockPutInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BytesMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *BytesMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockPutDone()
}
