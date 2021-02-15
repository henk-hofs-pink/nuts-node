// Code generated by MockGen. DO NOT EDIT.
// Source: core/echo.go

// Package core is a generated GoMock package.
package core

import (
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	reflect "reflect"
)

// MockEchoServer is a mock of EchoServer interface
type MockEchoServer struct {
	ctrl     *gomock.Controller
	recorder *MockEchoServerMockRecorder
}

// MockEchoServerMockRecorder is the mock recorder for MockEchoServer
type MockEchoServerMockRecorder struct {
	mock *MockEchoServer
}

// NewMockEchoServer creates a new mock instance
func NewMockEchoServer(ctrl *gomock.Controller) *MockEchoServer {
	mock := &MockEchoServer{ctrl: ctrl}
	mock.recorder = &MockEchoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEchoServer) EXPECT() *MockEchoServerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockEchoServer) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	m.ctrl.T.Helper()
	varargs := []interface{}{method, path, handler}
	for _, a := range middleware {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*echo.Route)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockEchoServerMockRecorder) Add(method, path, handler interface{}, middleware ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{method, path, handler}, middleware...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEchoServer)(nil).Add), varargs...)
}

// Start mocks base method
func (m *MockEchoServer) Start(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockEchoServerMockRecorder) Start(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEchoServer)(nil).Start), address)
}

// MockEchoRouter is a mock of EchoRouter interface
type MockEchoRouter struct {
	ctrl     *gomock.Controller
	recorder *MockEchoRouterMockRecorder
}

// MockEchoRouterMockRecorder is the mock recorder for MockEchoRouter
type MockEchoRouterMockRecorder struct {
	mock *MockEchoRouter
}

// NewMockEchoRouter creates a new mock instance
func NewMockEchoRouter(ctrl *gomock.Controller) *MockEchoRouter {
	mock := &MockEchoRouter{ctrl: ctrl}
	mock.recorder = &MockEchoRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEchoRouter) EXPECT() *MockEchoRouterMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockEchoRouter) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	m.ctrl.T.Helper()
	varargs := []interface{}{method, path, handler}
	for _, a := range middleware {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*echo.Route)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockEchoRouterMockRecorder) Add(method, path, handler interface{}, middleware ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{method, path, handler}, middleware...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEchoRouter)(nil).Add), varargs...)
}