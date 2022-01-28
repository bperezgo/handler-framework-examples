package main

import handlerLib "github.com/bperezgo/handler-framework"

type IDummyObject interface {
	RunSomething(req *handlerLib.HandlerRequest) string
}

type ConcreteDummyObject struct{}

func (ConcreteDummyObject) RunSomething(req *handlerLib.HandlerRequest) string {
	return "Ran something amazing"
}

// Mock used to replace the behavior in the test, thanks to the dependency injection
type MockDummyObject struct{}

func (do *MockDummyObject) RunSomething(req *handlerLib.HandlerRequest) string {
	return "Ran another amazing thing from mock of dummyObject"
}
