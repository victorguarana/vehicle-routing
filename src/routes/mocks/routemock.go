// Code generated by MockGen. DO NOT EDIT.
// Source: src/routes/route.go
//
// Generated by this command:
//
//	mockgen -source=src/routes/route.go -destination=src/routes/mocks/routemock.go -package=mockroutes
//
// Package mockroutes is a generated GoMock package.
package mockroutes

import (
	reflect "reflect"

	gps "github.com/victorguarana/go-vehicle-route/src/gps"
	routes "github.com/victorguarana/go-vehicle-route/src/routes"
	vehicles "github.com/victorguarana/go-vehicle-route/src/vehicles"
	gomock "go.uber.org/mock/gomock"
)

// MockIRoute is a mock of IRoute interface.
type MockIRoute struct {
	ctrl     *gomock.Controller
	recorder *MockIRouteMockRecorder
}

// MockIRouteMockRecorder is the mock recorder for MockIRoute.
type MockIRouteMockRecorder struct {
	mock *MockIRoute
}

// NewMockIRoute creates a new mock instance.
func NewMockIRoute(ctrl *gomock.Controller) *MockIRoute {
	mock := &MockIRoute{ctrl: ctrl}
	mock.recorder = &MockIRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoute) EXPECT() *MockIRouteMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockIRoute) Append(arg0 *gps.Point) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Append indicates an expected call of Append.
func (mr *MockIRouteMockRecorder) Append(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockIRoute)(nil).Append), arg0)
}

// Car mocks base method.
func (m *MockIRoute) Car() vehicles.ICar {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Car")
	ret0, _ := ret[0].(vehicles.ICar)
	return ret0
}

// Car indicates an expected call of Car.
func (mr *MockIRouteMockRecorder) Car() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Car", reflect.TypeOf((*MockIRoute)(nil).Car))
}

// First mocks base method.
func (m *MockIRoute) First() routes.ICarStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(routes.ICarStop)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockIRouteMockRecorder) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockIRoute)(nil).First))
}

// Last mocks base method.
func (m *MockIRoute) Last() routes.ICarStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(routes.ICarStop)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockIRouteMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockIRoute)(nil).Last))
}
