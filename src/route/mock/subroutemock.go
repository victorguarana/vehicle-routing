// Code generated by MockGen. DO NOT EDIT.
// Source: src/route/subroute.go
//
// Generated by this command:
//
//	mockgen -source=src/route/subroute.go -destination=src/route/mock/subroutemock.go
//

// Package mock_route is a generated GoMock package.
package mock_route

import (
	reflect "reflect"

	route "github.com/victorguarana/vehicle-routing/src/route"
	slc "github.com/victorguarana/vehicle-routing/src/slc"
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
func (m *MockISubRoute) Append(arg0 route.ISubStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append.
func (mr *MockISubRouteMockRecorder) Append(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockISubRoute)(nil).Append), arg0)
}

// First mocks base method.
func (m *MockISubRoute) First() route.ISubStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(route.ISubStop)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockISubRouteMockRecorder) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockISubRoute)(nil).First))
}

// InsertAt mocks base method.
func (m *MockISubRoute) InsertAt(index int, iSubStop route.ISubStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertAt", index, iSubStop)
}

// InsertAt indicates an expected call of InsertAt.
func (mr *MockISubRouteMockRecorder) InsertAt(index, iSubStop any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAt", reflect.TypeOf((*MockISubRoute)(nil).InsertAt), index, iSubStop)
}

// Iterator mocks base method.
func (m *MockISubRoute) Iterator() slc.Iterator[route.ISubStop] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(slc.Iterator[route.ISubStop])
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockISubRouteMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockISubRoute)(nil).Iterator))
}

// Last mocks base method.
func (m *MockISubRoute) Last() route.ISubStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(route.ISubStop)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockISubRouteMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockISubRoute)(nil).Last))
}

// Length mocks base method.
func (m *MockISubRoute) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockISubRouteMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockISubRoute)(nil).Length))
}

// RemoveSubStop mocks base method.
func (m *MockISubRoute) RemoveSubStop(index int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveSubStop", index)
}

// RemoveSubStop indicates an expected call of RemoveSubStop.
func (mr *MockISubRouteMockRecorder) RemoveSubStop(index any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubStop", reflect.TypeOf((*MockISubRoute)(nil).RemoveSubStop), index)
}

// Return mocks base method.
func (m *MockISubRoute) Return(arg0 route.IMainStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Return", arg0)
}

// Return indicates an expected call of Return.
func (mr *MockISubRouteMockRecorder) Return(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Return", reflect.TypeOf((*MockISubRoute)(nil).Return), arg0)
}

// ReturningStop mocks base method.
func (m *MockISubRoute) ReturningStop() route.IMainStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturningStop")
	ret0, _ := ret[0].(route.IMainStop)
	return ret0
}

// ReturningStop indicates an expected call of ReturningStop.
func (mr *MockISubRouteMockRecorder) ReturningStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturningStop", reflect.TypeOf((*MockISubRoute)(nil).ReturningStop))
}

// StartingStop mocks base method.
func (m *MockISubRoute) StartingStop() route.IMainStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartingStop")
	ret0, _ := ret[0].(route.IMainStop)
	return ret0
}

// StartingStop indicates an expected call of StartingStop.
func (mr *MockISubRouteMockRecorder) StartingStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartingStop", reflect.TypeOf((*MockISubRoute)(nil).StartingStop))
}
