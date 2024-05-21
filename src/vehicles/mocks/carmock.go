// Code generated by MockGen. DO NOT EDIT.
// Source: src/vehicles/car.go
//
// Generated by this command:
//
//	mockgen -source=src/vehicles/car.go -destination=src/vehicles/mocks/carmock.go -package=mockvehicles
//

// Package mockvehicles is a generated GoMock package.
package mockvehicles

import (
	reflect "reflect"

	gps "github.com/victorguarana/vehicle-routing/src/gps"
	vehicles "github.com/victorguarana/vehicle-routing/src/vehicles"
	gomock "go.uber.org/mock/gomock"
)

// MockICar is a mock of ICar interface.
type MockICar struct {
	ctrl     *gomock.Controller
	recorder *MockICarMockRecorder
}

// MockICarMockRecorder is the mock recorder for MockICar.
type MockICarMockRecorder struct {
	mock *MockICar
}

// NewMockICar creates a new mock instance.
func NewMockICar(ctrl *gomock.Controller) *MockICar {
	mock := &MockICar{ctrl: ctrl}
	mock.recorder = &MockICarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICar) EXPECT() *MockICarMockRecorder {
	return m.recorder
}

// ActualPoint mocks base method.
func (m *MockICar) ActualPoint() gps.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActualPoint")
	ret0, _ := ret[0].(gps.Point)
	return ret0
}

// ActualPoint indicates an expected call of ActualPoint.
func (mr *MockICarMockRecorder) ActualPoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActualPoint", reflect.TypeOf((*MockICar)(nil).ActualPoint))
}

// Drones mocks base method.
func (m *MockICar) Drones() []vehicles.IDrone {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drones")
	ret0, _ := ret[0].([]vehicles.IDrone)
	return ret0
}

// Drones indicates an expected call of Drones.
func (mr *MockICarMockRecorder) Drones() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drones", reflect.TypeOf((*MockICar)(nil).Drones))
}

// Efficiency mocks base method.
func (m *MockICar) Efficiency() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Efficiency")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Efficiency indicates an expected call of Efficiency.
func (mr *MockICarMockRecorder) Efficiency() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Efficiency", reflect.TypeOf((*MockICar)(nil).Efficiency))
}

// Move mocks base method.
func (m *MockICar) Move(destination gps.Point) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Move", destination)
}

// Move indicates an expected call of Move.
func (mr *MockICarMockRecorder) Move(destination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockICar)(nil).Move), destination)
}

// Name mocks base method.
func (m *MockICar) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockICarMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockICar)(nil).Name))
}

// NewDrone mocks base method.
func (m *MockICar) NewDrone(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewDrone", name)
}

// NewDrone indicates an expected call of NewDrone.
func (mr *MockICarMockRecorder) NewDrone(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDrone", reflect.TypeOf((*MockICar)(nil).NewDrone), name)
}

// Speed mocks base method.
func (m *MockICar) Speed() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Speed")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Speed indicates an expected call of Speed.
func (mr *MockICarMockRecorder) Speed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Speed", reflect.TypeOf((*MockICar)(nil).Speed))
}

// Support mocks base method.
func (m *MockICar) Support(arg0 ...gps.Point) bool {
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
func (mr *MockICarMockRecorder) Support(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Support", reflect.TypeOf((*MockICar)(nil).Support), arg0...)
}
