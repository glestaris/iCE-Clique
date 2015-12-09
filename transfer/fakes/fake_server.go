// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/glestaris/ice-clique/transfer"
)

type FakeServer struct {
	ServeStub         func()
	serveMutex        sync.RWMutex
	serveArgsForCall  []struct{}
	PauseStub         func()
	pauseMutex        sync.RWMutex
	pauseArgsForCall  []struct{}
	ResumeStub        func()
	resumeMutex       sync.RWMutex
	resumeArgsForCall []struct{}
	CloseStub         func() error
	closeMutex        sync.RWMutex
	closeArgsForCall  []struct{}
	closeReturns      struct {
		result1 error
	}
	LastTrasferStub        func() transfer.TransferResults
	lastTrasferMutex       sync.RWMutex
	lastTrasferArgsForCall []struct{}
	lastTrasferReturns     struct {
		result1 transfer.TransferResults
	}
}

func (fake *FakeServer) Serve() {
	fake.serveMutex.Lock()
	fake.serveArgsForCall = append(fake.serveArgsForCall, struct{}{})
	fake.serveMutex.Unlock()
	if fake.ServeStub != nil {
		fake.ServeStub()
	}
}

func (fake *FakeServer) ServeCallCount() int {
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	return len(fake.serveArgsForCall)
}

func (fake *FakeServer) Pause() {
	fake.pauseMutex.Lock()
	fake.pauseArgsForCall = append(fake.pauseArgsForCall, struct{}{})
	fake.pauseMutex.Unlock()
	if fake.PauseStub != nil {
		fake.PauseStub()
	}
}

func (fake *FakeServer) PauseCallCount() int {
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	return len(fake.pauseArgsForCall)
}

func (fake *FakeServer) Resume() {
	fake.resumeMutex.Lock()
	fake.resumeArgsForCall = append(fake.resumeArgsForCall, struct{}{})
	fake.resumeMutex.Unlock()
	if fake.ResumeStub != nil {
		fake.ResumeStub()
	}
}

func (fake *FakeServer) ResumeCallCount() int {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return len(fake.resumeArgsForCall)
}

func (fake *FakeServer) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeServer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeServer) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServer) LastTrasfer() transfer.TransferResults {
	fake.lastTrasferMutex.Lock()
	fake.lastTrasferArgsForCall = append(fake.lastTrasferArgsForCall, struct{}{})
	fake.lastTrasferMutex.Unlock()
	if fake.LastTrasferStub != nil {
		return fake.LastTrasferStub()
	} else {
		return fake.lastTrasferReturns.result1
	}
}

func (fake *FakeServer) LastTrasferCallCount() int {
	fake.lastTrasferMutex.RLock()
	defer fake.lastTrasferMutex.RUnlock()
	return len(fake.lastTrasferArgsForCall)
}

func (fake *FakeServer) LastTrasferReturns(result1 transfer.TransferResults) {
	fake.LastTrasferStub = nil
	fake.lastTrasferReturns = struct {
		result1 transfer.TransferResults
	}{result1}
}

var _ transfer.Server = new(FakeServer)
