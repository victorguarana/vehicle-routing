// Code generated by MockGen. DO NOT EDIT.
// Source: drone.go
//
// Generated by this command:
//
//	mockgen -source=drone.go -destination=mock/dronemock.go
//

// Package mock_vehicle is a generated GoMock package.
package mock_vehicle

import (
	reflect "reflect"

	gps "github.com/victorguarana/vehicle-routing/internal/gps"
	gomock "go.uber.org/mock/gomock"
)

// MockIDrone is a mock of IDrone interface.
type MockIDrone struct {
	ctrl     *gomock.Controller
	recorder *MockIDroneMockRecorder
}

// MockIDroneMockRecorder is the mock recorder for MockIDrone.
type MockIDroneMockRecorder struct {
	mock *MockIDrone
}

// NewMockIDrone creates a new mock instance.
func NewMockIDrone(ctrl *gomock.Controller) *MockIDrone {
	mock := &MockIDrone{ctrl: ctrl}
	mock.recorder = &MockIDroneMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDrone) EXPECT() *MockIDroneMockRecorder {
	return m.recorder
}

// ActualPoint mocks base method.
func (m *MockIDrone) ActualPoint() gps.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActualPoint")
	ret0, _ := ret[0].(gps.Point)
	return ret0
}

// ActualPoint indicates an expected call of ActualPoint.
func (mr *MockIDroneMockRecorder) ActualPoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActualPoint", reflect.TypeOf((*MockIDrone)(nil).ActualPoint))
}

// CanReach mocks base method.
func (m *MockIDrone) CanReach(arg0 ...gps.Point) bool {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanReach", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanReach indicates an expected call of CanReach.
func (mr *MockIDroneMockRecorder) CanReach(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanReach", reflect.TypeOf((*MockIDrone)(nil).CanReach), arg0...)
}

// Efficiency mocks base method.
func (m *MockIDrone) Efficiency() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Efficiency")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Efficiency indicates an expected call of Efficiency.
func (mr *MockIDroneMockRecorder) Efficiency() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Efficiency", reflect.TypeOf((*MockIDrone)(nil).Efficiency))
}

// IsFlying mocks base method.
func (m *MockIDrone) IsFlying() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFlying")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFlying indicates an expected call of IsFlying.
func (mr *MockIDroneMockRecorder) IsFlying() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFlying", reflect.TypeOf((*MockIDrone)(nil).IsFlying))
}

// Land mocks base method.
func (m *MockIDrone) Land(destination gps.Point) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Land", destination)
}

// Land indicates an expected call of Land.
func (mr *MockIDroneMockRecorder) Land(destination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Land", reflect.TypeOf((*MockIDrone)(nil).Land), destination)
}

// Move mocks base method.
func (m *MockIDrone) Move(destination gps.Point) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Move", destination)
}

// Move indicates an expected call of Move.
func (mr *MockIDroneMockRecorder) Move(destination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockIDrone)(nil).Move), destination)
}

// Name mocks base method.
func (m *MockIDrone) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockIDroneMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockIDrone)(nil).Name))
}

// Speed mocks base method.
func (m *MockIDrone) Speed() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Speed")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Speed indicates an expected call of Speed.
func (mr *MockIDroneMockRecorder) Speed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Speed", reflect.TypeOf((*MockIDrone)(nil).Speed))
}

// Support mocks base method.
func (m *MockIDrone) Support(arg0 ...gps.Point) bool {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Support", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Support indicates an expected call of Support.
func (mr *MockIDroneMockRecorder) Support(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Support", reflect.TypeOf((*MockIDrone)(nil).Support), arg0...)
}

// TakeOff mocks base method.
func (m *MockIDrone) TakeOff() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TakeOff")
}

// TakeOff indicates an expected call of TakeOff.
func (mr *MockIDroneMockRecorder) TakeOff() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeOff", reflect.TypeOf((*MockIDrone)(nil).TakeOff))
}