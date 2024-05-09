// Code generated by MockGen. DO NOT EDIT.
// Source: src/routes/subroute.go
//
// Generated by this command:
//
//	mockgen -source=src/routes/subroute.go -destination=src/routes/mocks/subroute.go -package=mockroutes
//

// Package mockroutes is a generated GoMock package.
package mockroutes

import (
	reflect "reflect"

	routes "github.com/victorguarana/go-vehicle-route/src/routes"
	gomock "go.uber.org/mock/gomock"
)

// MockISubRoute is a mock of ISubRoute interface.
type MockISubRoute struct {
	ctrl     *gomock.Controller
	recorder *MockISubRouteMockRecorder
}

// MockISubRouteMockRecorder is the mock recorder for MockISubRoute.
type MockISubRouteMockRecorder struct {
	mock *MockISubRoute
}

// NewMockISubRoute creates a new mock instance.
func NewMockISubRoute(ctrl *gomock.Controller) *MockISubRoute {
	mock := &MockISubRoute{ctrl: ctrl}
	mock.recorder = &MockISubRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubRoute) EXPECT() *MockISubRouteMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockISubRoute) Append(arg0 routes.ISubStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append.
func (mr *MockISubRouteMockRecorder) Append(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockISubRoute)(nil).Append), arg0)
}

// Return mocks base method.
func (m *MockISubRoute) Return(arg0 routes.IMainStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Return", arg0)
}

// Return indicates an expected call of Return.
func (mr *MockISubRouteMockRecorder) Return(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Return", reflect.TypeOf((*MockISubRoute)(nil).Return), arg0)
}

// ReturningPoint mocks base method.
func (m *MockISubRoute) ReturningPoint() routes.IMainStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturningPoint")
	ret0, _ := ret[0].(routes.IMainStop)
	return ret0
}

// ReturningPoint indicates an expected call of ReturningPoint.
func (mr *MockISubRouteMockRecorder) ReturningPoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturningPoint", reflect.TypeOf((*MockISubRoute)(nil).ReturningPoint))
}

// StartingPoint mocks base method.
func (m *MockISubRoute) StartingPoint() routes.IMainStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartingPoint")
	ret0, _ := ret[0].(routes.IMainStop)
	return ret0
}

// StartingPoint indicates an expected call of StartingPoint.
func (mr *MockISubRouteMockRecorder) StartingPoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartingPoint", reflect.TypeOf((*MockISubRoute)(nil).StartingPoint))
}
