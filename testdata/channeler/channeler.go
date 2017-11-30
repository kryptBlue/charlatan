// generated by "charlatan -dir=testdata/channeler -output=testdata/channeler/channeler.go Channeler".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// ChannelInvocation represents a single call of FakeChanneler.Channel
type ChannelInvocation struct {
	Parameters struct {
		Ident1 chan int
	}
	Results struct {
		Ident2 chan int
	}
}

// ChannelReceiveInvocation represents a single call of FakeChanneler.ChannelReceive
type ChannelReceiveInvocation struct {
	Parameters struct {
		Ident3 <-chan int
	}
	Results struct {
		Ident4 <-chan int
	}
}

// ChannelSendInvocation represents a single call of FakeChanneler.ChannelSend
type ChannelSendInvocation struct {
	Parameters struct {
		Ident5 chan<- int
	}
	Results struct {
		Ident6 chan<- int
	}
}

// ChannelPointerInvocation represents a single call of FakeChanneler.ChannelPointer
type ChannelPointerInvocation struct {
	Parameters struct {
		Ident7 *chan int
	}
	Results struct {
		Ident8 *chan int
	}
}

// ChannelInterfaceInvocation represents a single call of FakeChanneler.ChannelInterface
type ChannelInterfaceInvocation struct {
	Parameters struct {
		Ident9 chan interface{}
	}
	Results struct {
		Ident10 chan interface{}
	}
}

/*
FakeChanneler is a mock implementation of Channeler for testing.
Use it in your tests as in this example:

	package example

	func TestWithChanneler(t *testing.T) {
		f := &main.FakeChanneler{
			ChannelHook: func(ident1 chan int) (ident2 chan int) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeChannel ...
		f.AssertChannelCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeChannel.
*/
type FakeChanneler struct {
	ChannelHook          func(chan int) chan int
	ChannelReceiveHook   func(<-chan int) <-chan int
	ChannelSendHook      func(chan<- int) chan<- int
	ChannelPointerHook   func(*chan int) *chan int
	ChannelInterfaceHook func(chan interface{}) chan interface{}

	ChannelCalls          []*ChannelInvocation
	ChannelReceiveCalls   []*ChannelReceiveInvocation
	ChannelSendCalls      []*ChannelSendInvocation
	ChannelPointerCalls   []*ChannelPointerInvocation
	ChannelInterfaceCalls []*ChannelInterfaceInvocation
}

// NewFakeChannelerDefaultPanic returns an instance of FakeChanneler with all hooks configured to panic
func NewFakeChannelerDefaultPanic() *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			panic("Unexpected call to Channeler.Channel")
		},
		ChannelReceiveHook: func(<-chan int) (ident4 <-chan int) {
			panic("Unexpected call to Channeler.ChannelReceive")
		},
		ChannelSendHook: func(chan<- int) (ident6 chan<- int) {
			panic("Unexpected call to Channeler.ChannelSend")
		},
		ChannelPointerHook: func(*chan int) (ident8 *chan int) {
			panic("Unexpected call to Channeler.ChannelPointer")
		},
		ChannelInterfaceHook: func(chan interface{}) (ident10 chan interface{}) {
			panic("Unexpected call to Channeler.ChannelInterface")
		},
	}
}

// NewFakeChannelerDefaultFatal returns an instance of FakeChanneler with all hooks configured to call t.Fatal
func NewFakeChannelerDefaultFatal(t *testing.T) *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			t.Fatal("Unexpected call to Channeler.Channel")
			return
		},
		ChannelReceiveHook: func(<-chan int) (ident4 <-chan int) {
			t.Fatal("Unexpected call to Channeler.ChannelReceive")
			return
		},
		ChannelSendHook: func(chan<- int) (ident6 chan<- int) {
			t.Fatal("Unexpected call to Channeler.ChannelSend")
			return
		},
		ChannelPointerHook: func(*chan int) (ident8 *chan int) {
			t.Fatal("Unexpected call to Channeler.ChannelPointer")
			return
		},
		ChannelInterfaceHook: func(chan interface{}) (ident10 chan interface{}) {
			t.Fatal("Unexpected call to Channeler.ChannelInterface")
			return
		},
	}
}

// NewFakeChannelerDefaultError returns an instance of FakeChanneler with all hooks configured to call t.Error
func NewFakeChannelerDefaultError(t *testing.T) *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			t.Error("Unexpected call to Channeler.Channel")
			return
		},
		ChannelReceiveHook: func(<-chan int) (ident4 <-chan int) {
			t.Error("Unexpected call to Channeler.ChannelReceive")
			return
		},
		ChannelSendHook: func(chan<- int) (ident6 chan<- int) {
			t.Error("Unexpected call to Channeler.ChannelSend")
			return
		},
		ChannelPointerHook: func(*chan int) (ident8 *chan int) {
			t.Error("Unexpected call to Channeler.ChannelPointer")
			return
		},
		ChannelInterfaceHook: func(chan interface{}) (ident10 chan interface{}) {
			t.Error("Unexpected call to Channeler.ChannelInterface")
			return
		},
	}
}

func (_f1 *FakeChanneler) Channel(ident1 chan int) (ident2 chan int) {
	invocation := new(ChannelInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f1.ChannelHook(ident1)

	invocation.Results.Ident2 = ident2

	_f1.ChannelCalls = append(_f1.ChannelCalls, invocation)

	return
}

// ChannelCalled returns true if FakeChanneler.Channel was called
func (f *FakeChanneler) ChannelCalled() bool {
	return len(f.ChannelCalls) != 0
}

// AssertChannelCalled calls t.Error if FakeChanneler.Channel was not called
func (f *FakeChanneler) AssertChannelCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelCalls) == 0 {
		t.Error("FakeChanneler.Channel not called, expected at least one")
	}
}

// ChannelNotCalled returns true if FakeChanneler.Channel was not called
func (f *FakeChanneler) ChannelNotCalled() bool {
	return len(f.ChannelCalls) == 0
}

// AssertChannelNotCalled calls t.Error if FakeChanneler.Channel was called
func (f *FakeChanneler) AssertChannelNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelCalls) != 0 {
		t.Error("FakeChanneler.Channel called, expected none")
	}
}

// ChannelCalledOnce returns true if FakeChanneler.Channel was called exactly once
func (f *FakeChanneler) ChannelCalledOnce() bool {
	return len(f.ChannelCalls) == 1
}

// AssertChannelCalledOnce calls t.Error if FakeChanneler.Channel was not called exactly once
func (f *FakeChanneler) AssertChannelCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ChannelCalls) != 1 {
		t.Errorf("FakeChanneler.Channel called %d times, expected 1", len(f.ChannelCalls))
	}
}

// ChannelCalledN returns true if FakeChanneler.Channel was called at least n times
func (f *FakeChanneler) ChannelCalledN(n int) bool {
	return len(f.ChannelCalls) >= n
}

// AssertChannelCalledN calls t.Error if FakeChanneler.Channel was called less than n times
func (f *FakeChanneler) AssertChannelCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ChannelCalls) < n {
		t.Errorf("FakeChanneler.Channel called %d times, expected >= %d", len(f.ChannelCalls), n)
	}
}

// ChannelCalledWith returns true if FakeChanneler.Channel was called with the given values
func (_f2 *FakeChanneler) ChannelCalledWith(ident1 chan int) (found bool) {
	for _, call := range _f2.ChannelCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertChannelCalledWith calls t.Error if FakeChanneler.Channel was not called with the given values
func (_f3 *FakeChanneler) AssertChannelCalledWith(t *testing.T, ident1 chan int) {
	t.Helper()
	var found bool
	for _, call := range _f3.ChannelCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeChanneler.Channel not called with expected parameters")
	}
}

// ChannelCalledOnceWith returns true if FakeChanneler.Channel was called exactly once with the given values
func (_f4 *FakeChanneler) ChannelCalledOnceWith(ident1 chan int) bool {
	var count int
	for _, call := range _f4.ChannelCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertChannelCalledOnceWith calls t.Error if FakeChanneler.Channel was not called exactly once with the given values
func (_f5 *FakeChanneler) AssertChannelCalledOnceWith(t *testing.T, ident1 chan int) {
	t.Helper()
	var count int
	for _, call := range _f5.ChannelCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeChanneler.Channel called %d times with expected parameters, expected one", count)
	}
}

// ChannelResultsForCall returns the result values for the first call to FakeChanneler.Channel with the given values
func (_f6 *FakeChanneler) ChannelResultsForCall(ident1 chan int) (ident2 chan int, found bool) {
	for _, call := range _f6.ChannelCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeChanneler) ChannelReceive(ident3 <-chan int) (ident4 <-chan int) {
	invocation := new(ChannelReceiveInvocation)

	invocation.Parameters.Ident3 = ident3

	ident4 = _f7.ChannelReceiveHook(ident3)

	invocation.Results.Ident4 = ident4

	_f7.ChannelReceiveCalls = append(_f7.ChannelReceiveCalls, invocation)

	return
}

// ChannelReceiveCalled returns true if FakeChanneler.ChannelReceive was called
func (f *FakeChanneler) ChannelReceiveCalled() bool {
	return len(f.ChannelReceiveCalls) != 0
}

// AssertChannelReceiveCalled calls t.Error if FakeChanneler.ChannelReceive was not called
func (f *FakeChanneler) AssertChannelReceiveCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelReceiveCalls) == 0 {
		t.Error("FakeChanneler.ChannelReceive not called, expected at least one")
	}
}

// ChannelReceiveNotCalled returns true if FakeChanneler.ChannelReceive was not called
func (f *FakeChanneler) ChannelReceiveNotCalled() bool {
	return len(f.ChannelReceiveCalls) == 0
}

// AssertChannelReceiveNotCalled calls t.Error if FakeChanneler.ChannelReceive was called
func (f *FakeChanneler) AssertChannelReceiveNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelReceiveCalls) != 0 {
		t.Error("FakeChanneler.ChannelReceive called, expected none")
	}
}

// ChannelReceiveCalledOnce returns true if FakeChanneler.ChannelReceive was called exactly once
func (f *FakeChanneler) ChannelReceiveCalledOnce() bool {
	return len(f.ChannelReceiveCalls) == 1
}

// AssertChannelReceiveCalledOnce calls t.Error if FakeChanneler.ChannelReceive was not called exactly once
func (f *FakeChanneler) AssertChannelReceiveCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ChannelReceiveCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelReceive called %d times, expected 1", len(f.ChannelReceiveCalls))
	}
}

// ChannelReceiveCalledN returns true if FakeChanneler.ChannelReceive was called at least n times
func (f *FakeChanneler) ChannelReceiveCalledN(n int) bool {
	return len(f.ChannelReceiveCalls) >= n
}

// AssertChannelReceiveCalledN calls t.Error if FakeChanneler.ChannelReceive was called less than n times
func (f *FakeChanneler) AssertChannelReceiveCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ChannelReceiveCalls) < n {
		t.Errorf("FakeChanneler.ChannelReceive called %d times, expected >= %d", len(f.ChannelReceiveCalls), n)
	}
}

// ChannelReceiveCalledWith returns true if FakeChanneler.ChannelReceive was called with the given values
func (_f8 *FakeChanneler) ChannelReceiveCalledWith(ident3 <-chan int) (found bool) {
	for _, call := range _f8.ChannelReceiveCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			found = true
			break
		}
	}

	return
}

// AssertChannelReceiveCalledWith calls t.Error if FakeChanneler.ChannelReceive was not called with the given values
func (_f9 *FakeChanneler) AssertChannelReceiveCalledWith(t *testing.T, ident3 <-chan int) {
	t.Helper()
	var found bool
	for _, call := range _f9.ChannelReceiveCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeChanneler.ChannelReceive not called with expected parameters")
	}
}

// ChannelReceiveCalledOnceWith returns true if FakeChanneler.ChannelReceive was called exactly once with the given values
func (_f10 *FakeChanneler) ChannelReceiveCalledOnceWith(ident3 <-chan int) bool {
	var count int
	for _, call := range _f10.ChannelReceiveCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			count++
		}
	}

	return count == 1
}

// AssertChannelReceiveCalledOnceWith calls t.Error if FakeChanneler.ChannelReceive was not called exactly once with the given values
func (_f11 *FakeChanneler) AssertChannelReceiveCalledOnceWith(t *testing.T, ident3 <-chan int) {
	t.Helper()
	var count int
	for _, call := range _f11.ChannelReceiveCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeChanneler.ChannelReceive called %d times with expected parameters, expected one", count)
	}
}

// ChannelReceiveResultsForCall returns the result values for the first call to FakeChanneler.ChannelReceive with the given values
func (_f12 *FakeChanneler) ChannelReceiveResultsForCall(ident3 <-chan int) (ident4 <-chan int, found bool) {
	for _, call := range _f12.ChannelReceiveCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}

func (_f13 *FakeChanneler) ChannelSend(ident5 chan<- int) (ident6 chan<- int) {
	invocation := new(ChannelSendInvocation)

	invocation.Parameters.Ident5 = ident5

	ident6 = _f13.ChannelSendHook(ident5)

	invocation.Results.Ident6 = ident6

	_f13.ChannelSendCalls = append(_f13.ChannelSendCalls, invocation)

	return
}

// ChannelSendCalled returns true if FakeChanneler.ChannelSend was called
func (f *FakeChanneler) ChannelSendCalled() bool {
	return len(f.ChannelSendCalls) != 0
}

// AssertChannelSendCalled calls t.Error if FakeChanneler.ChannelSend was not called
func (f *FakeChanneler) AssertChannelSendCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelSendCalls) == 0 {
		t.Error("FakeChanneler.ChannelSend not called, expected at least one")
	}
}

// ChannelSendNotCalled returns true if FakeChanneler.ChannelSend was not called
func (f *FakeChanneler) ChannelSendNotCalled() bool {
	return len(f.ChannelSendCalls) == 0
}

// AssertChannelSendNotCalled calls t.Error if FakeChanneler.ChannelSend was called
func (f *FakeChanneler) AssertChannelSendNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelSendCalls) != 0 {
		t.Error("FakeChanneler.ChannelSend called, expected none")
	}
}

// ChannelSendCalledOnce returns true if FakeChanneler.ChannelSend was called exactly once
func (f *FakeChanneler) ChannelSendCalledOnce() bool {
	return len(f.ChannelSendCalls) == 1
}

// AssertChannelSendCalledOnce calls t.Error if FakeChanneler.ChannelSend was not called exactly once
func (f *FakeChanneler) AssertChannelSendCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ChannelSendCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelSend called %d times, expected 1", len(f.ChannelSendCalls))
	}
}

// ChannelSendCalledN returns true if FakeChanneler.ChannelSend was called at least n times
func (f *FakeChanneler) ChannelSendCalledN(n int) bool {
	return len(f.ChannelSendCalls) >= n
}

// AssertChannelSendCalledN calls t.Error if FakeChanneler.ChannelSend was called less than n times
func (f *FakeChanneler) AssertChannelSendCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ChannelSendCalls) < n {
		t.Errorf("FakeChanneler.ChannelSend called %d times, expected >= %d", len(f.ChannelSendCalls), n)
	}
}

// ChannelSendCalledWith returns true if FakeChanneler.ChannelSend was called with the given values
func (_f14 *FakeChanneler) ChannelSendCalledWith(ident5 chan<- int) (found bool) {
	for _, call := range _f14.ChannelSendCalls {
		if reflect.DeepEqual(call.Parameters.Ident5, ident5) {
			found = true
			break
		}
	}

	return
}

// AssertChannelSendCalledWith calls t.Error if FakeChanneler.ChannelSend was not called with the given values
func (_f15 *FakeChanneler) AssertChannelSendCalledWith(t *testing.T, ident5 chan<- int) {
	t.Helper()
	var found bool
	for _, call := range _f15.ChannelSendCalls {
		if reflect.DeepEqual(call.Parameters.Ident5, ident5) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeChanneler.ChannelSend not called with expected parameters")
	}
}

// ChannelSendCalledOnceWith returns true if FakeChanneler.ChannelSend was called exactly once with the given values
func (_f16 *FakeChanneler) ChannelSendCalledOnceWith(ident5 chan<- int) bool {
	var count int
	for _, call := range _f16.ChannelSendCalls {
		if reflect.DeepEqual(call.Parameters.Ident5, ident5) {
			count++
		}
	}

	return count == 1
}

// AssertChannelSendCalledOnceWith calls t.Error if FakeChanneler.ChannelSend was not called exactly once with the given values
func (_f17 *FakeChanneler) AssertChannelSendCalledOnceWith(t *testing.T, ident5 chan<- int) {
	t.Helper()
	var count int
	for _, call := range _f17.ChannelSendCalls {
		if reflect.DeepEqual(call.Parameters.Ident5, ident5) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeChanneler.ChannelSend called %d times with expected parameters, expected one", count)
	}
}

// ChannelSendResultsForCall returns the result values for the first call to FakeChanneler.ChannelSend with the given values
func (_f18 *FakeChanneler) ChannelSendResultsForCall(ident5 chan<- int) (ident6 chan<- int, found bool) {
	for _, call := range _f18.ChannelSendCalls {
		if reflect.DeepEqual(call.Parameters.Ident5, ident5) {
			ident6 = call.Results.Ident6
			found = true
			break
		}
	}

	return
}

func (_f19 *FakeChanneler) ChannelPointer(ident7 *chan int) (ident8 *chan int) {
	invocation := new(ChannelPointerInvocation)

	invocation.Parameters.Ident7 = ident7

	ident8 = _f19.ChannelPointerHook(ident7)

	invocation.Results.Ident8 = ident8

	_f19.ChannelPointerCalls = append(_f19.ChannelPointerCalls, invocation)

	return
}

// ChannelPointerCalled returns true if FakeChanneler.ChannelPointer was called
func (f *FakeChanneler) ChannelPointerCalled() bool {
	return len(f.ChannelPointerCalls) != 0
}

// AssertChannelPointerCalled calls t.Error if FakeChanneler.ChannelPointer was not called
func (f *FakeChanneler) AssertChannelPointerCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelPointerCalls) == 0 {
		t.Error("FakeChanneler.ChannelPointer not called, expected at least one")
	}
}

// ChannelPointerNotCalled returns true if FakeChanneler.ChannelPointer was not called
func (f *FakeChanneler) ChannelPointerNotCalled() bool {
	return len(f.ChannelPointerCalls) == 0
}

// AssertChannelPointerNotCalled calls t.Error if FakeChanneler.ChannelPointer was called
func (f *FakeChanneler) AssertChannelPointerNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelPointerCalls) != 0 {
		t.Error("FakeChanneler.ChannelPointer called, expected none")
	}
}

// ChannelPointerCalledOnce returns true if FakeChanneler.ChannelPointer was called exactly once
func (f *FakeChanneler) ChannelPointerCalledOnce() bool {
	return len(f.ChannelPointerCalls) == 1
}

// AssertChannelPointerCalledOnce calls t.Error if FakeChanneler.ChannelPointer was not called exactly once
func (f *FakeChanneler) AssertChannelPointerCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ChannelPointerCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelPointer called %d times, expected 1", len(f.ChannelPointerCalls))
	}
}

// ChannelPointerCalledN returns true if FakeChanneler.ChannelPointer was called at least n times
func (f *FakeChanneler) ChannelPointerCalledN(n int) bool {
	return len(f.ChannelPointerCalls) >= n
}

// AssertChannelPointerCalledN calls t.Error if FakeChanneler.ChannelPointer was called less than n times
func (f *FakeChanneler) AssertChannelPointerCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ChannelPointerCalls) < n {
		t.Errorf("FakeChanneler.ChannelPointer called %d times, expected >= %d", len(f.ChannelPointerCalls), n)
	}
}

// ChannelPointerCalledWith returns true if FakeChanneler.ChannelPointer was called with the given values
func (_f20 *FakeChanneler) ChannelPointerCalledWith(ident7 *chan int) (found bool) {
	for _, call := range _f20.ChannelPointerCalls {
		if reflect.DeepEqual(call.Parameters.Ident7, ident7) {
			found = true
			break
		}
	}

	return
}

// AssertChannelPointerCalledWith calls t.Error if FakeChanneler.ChannelPointer was not called with the given values
func (_f21 *FakeChanneler) AssertChannelPointerCalledWith(t *testing.T, ident7 *chan int) {
	t.Helper()
	var found bool
	for _, call := range _f21.ChannelPointerCalls {
		if reflect.DeepEqual(call.Parameters.Ident7, ident7) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeChanneler.ChannelPointer not called with expected parameters")
	}
}

// ChannelPointerCalledOnceWith returns true if FakeChanneler.ChannelPointer was called exactly once with the given values
func (_f22 *FakeChanneler) ChannelPointerCalledOnceWith(ident7 *chan int) bool {
	var count int
	for _, call := range _f22.ChannelPointerCalls {
		if reflect.DeepEqual(call.Parameters.Ident7, ident7) {
			count++
		}
	}

	return count == 1
}

// AssertChannelPointerCalledOnceWith calls t.Error if FakeChanneler.ChannelPointer was not called exactly once with the given values
func (_f23 *FakeChanneler) AssertChannelPointerCalledOnceWith(t *testing.T, ident7 *chan int) {
	t.Helper()
	var count int
	for _, call := range _f23.ChannelPointerCalls {
		if reflect.DeepEqual(call.Parameters.Ident7, ident7) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeChanneler.ChannelPointer called %d times with expected parameters, expected one", count)
	}
}

// ChannelPointerResultsForCall returns the result values for the first call to FakeChanneler.ChannelPointer with the given values
func (_f24 *FakeChanneler) ChannelPointerResultsForCall(ident7 *chan int) (ident8 *chan int, found bool) {
	for _, call := range _f24.ChannelPointerCalls {
		if reflect.DeepEqual(call.Parameters.Ident7, ident7) {
			ident8 = call.Results.Ident8
			found = true
			break
		}
	}

	return
}

func (_f25 *FakeChanneler) ChannelInterface(ident9 chan interface{}) (ident10 chan interface{}) {
	invocation := new(ChannelInterfaceInvocation)

	invocation.Parameters.Ident9 = ident9

	ident10 = _f25.ChannelInterfaceHook(ident9)

	invocation.Results.Ident10 = ident10

	_f25.ChannelInterfaceCalls = append(_f25.ChannelInterfaceCalls, invocation)

	return
}

// ChannelInterfaceCalled returns true if FakeChanneler.ChannelInterface was called
func (f *FakeChanneler) ChannelInterfaceCalled() bool {
	return len(f.ChannelInterfaceCalls) != 0
}

// AssertChannelInterfaceCalled calls t.Error if FakeChanneler.ChannelInterface was not called
func (f *FakeChanneler) AssertChannelInterfaceCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) == 0 {
		t.Error("FakeChanneler.ChannelInterface not called, expected at least one")
	}
}

// ChannelInterfaceNotCalled returns true if FakeChanneler.ChannelInterface was not called
func (f *FakeChanneler) ChannelInterfaceNotCalled() bool {
	return len(f.ChannelInterfaceCalls) == 0
}

// AssertChannelInterfaceNotCalled calls t.Error if FakeChanneler.ChannelInterface was called
func (f *FakeChanneler) AssertChannelInterfaceNotCalled(t *testing.T) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) != 0 {
		t.Error("FakeChanneler.ChannelInterface called, expected none")
	}
}

// ChannelInterfaceCalledOnce returns true if FakeChanneler.ChannelInterface was called exactly once
func (f *FakeChanneler) ChannelInterfaceCalledOnce() bool {
	return len(f.ChannelInterfaceCalls) == 1
}

// AssertChannelInterfaceCalledOnce calls t.Error if FakeChanneler.ChannelInterface was not called exactly once
func (f *FakeChanneler) AssertChannelInterfaceCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelInterface called %d times, expected 1", len(f.ChannelInterfaceCalls))
	}
}

// ChannelInterfaceCalledN returns true if FakeChanneler.ChannelInterface was called at least n times
func (f *FakeChanneler) ChannelInterfaceCalledN(n int) bool {
	return len(f.ChannelInterfaceCalls) >= n
}

// AssertChannelInterfaceCalledN calls t.Error if FakeChanneler.ChannelInterface was called less than n times
func (f *FakeChanneler) AssertChannelInterfaceCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) < n {
		t.Errorf("FakeChanneler.ChannelInterface called %d times, expected >= %d", len(f.ChannelInterfaceCalls), n)
	}
}

// ChannelInterfaceCalledWith returns true if FakeChanneler.ChannelInterface was called with the given values
func (_f26 *FakeChanneler) ChannelInterfaceCalledWith(ident9 chan interface{}) (found bool) {
	for _, call := range _f26.ChannelInterfaceCalls {
		if reflect.DeepEqual(call.Parameters.Ident9, ident9) {
			found = true
			break
		}
	}

	return
}

// AssertChannelInterfaceCalledWith calls t.Error if FakeChanneler.ChannelInterface was not called with the given values
func (_f27 *FakeChanneler) AssertChannelInterfaceCalledWith(t *testing.T, ident9 chan interface{}) {
	t.Helper()
	var found bool
	for _, call := range _f27.ChannelInterfaceCalls {
		if reflect.DeepEqual(call.Parameters.Ident9, ident9) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeChanneler.ChannelInterface not called with expected parameters")
	}
}

// ChannelInterfaceCalledOnceWith returns true if FakeChanneler.ChannelInterface was called exactly once with the given values
func (_f28 *FakeChanneler) ChannelInterfaceCalledOnceWith(ident9 chan interface{}) bool {
	var count int
	for _, call := range _f28.ChannelInterfaceCalls {
		if reflect.DeepEqual(call.Parameters.Ident9, ident9) {
			count++
		}
	}

	return count == 1
}

// AssertChannelInterfaceCalledOnceWith calls t.Error if FakeChanneler.ChannelInterface was not called exactly once with the given values
func (_f29 *FakeChanneler) AssertChannelInterfaceCalledOnceWith(t *testing.T, ident9 chan interface{}) {
	t.Helper()
	var count int
	for _, call := range _f29.ChannelInterfaceCalls {
		if reflect.DeepEqual(call.Parameters.Ident9, ident9) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeChanneler.ChannelInterface called %d times with expected parameters, expected one", count)
	}
}

// ChannelInterfaceResultsForCall returns the result values for the first call to FakeChanneler.ChannelInterface with the given values
func (_f30 *FakeChanneler) ChannelInterfaceResultsForCall(ident9 chan interface{}) (ident10 chan interface{}, found bool) {
	for _, call := range _f30.ChannelInterfaceCalls {
		if reflect.DeepEqual(call.Parameters.Ident9, ident9) {
			ident10 = call.Results.Ident10
			found = true
			break
		}
	}

	return
}
