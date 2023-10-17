// Code generated by MockGen. DO NOT EDIT.
// Source: routes/route.go

// Package mock_routes is a generated GoMock package.
package mockroutes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gps "github.com/victorguarana/go-vehicle-route/src/gps"
	vehicles "github.com/victorguarana/go-vehicle-route/src/vehicles"
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
func (mr *MockIRouteMockRecorder) Append(arg0 interface{}) *gomock.Call {
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
func (m *MockIRoute) First() *gps.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(*gps.Point)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockIRouteMockRecorder) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockIRoute)(nil).First))
}

// Last mocks base method.
func (m *MockIRoute) Last() *gps.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(*gps.Point)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockIRouteMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockIRoute)(nil).Last))
}

// String mocks base method.
func (m *MockIRoute) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockIRouteMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockIRoute)(nil).String))
}
