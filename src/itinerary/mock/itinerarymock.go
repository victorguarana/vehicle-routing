// Code generated by MockGen. DO NOT EDIT.
// Source: src/itinerary/itinerary.go
//
// Generated by this command:
//
//	mockgen -source=src/itinerary/itinerary.go -destination=src/itinerary/mock/itinerarymock.go
//

// Package mock_itinerary is a generated GoMock package.
package mock_itinerary

import (
	reflect "reflect"

	itinerary "github.com/victorguarana/vehicle-routing/src/itinerary"
	gomock "go.uber.org/mock/gomock"
)

// MockItinerary is a mock of Itinerary interface.
type MockItinerary struct {
	ctrl     *gomock.Controller
	recorder *MockItineraryMockRecorder
}

// MockItineraryMockRecorder is the mock recorder for MockItinerary.
type MockItineraryMockRecorder struct {
	mock *MockItinerary
}

// NewMockItinerary creates a new mock instance.
func NewMockItinerary(ctrl *gomock.Controller) *MockItinerary {
	mock := &MockItinerary{ctrl: ctrl}
	mock.recorder = &MockItineraryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItinerary) EXPECT() *MockItineraryMockRecorder {
	return m.recorder
}

// Constructor mocks base method.
func (m *MockItinerary) Constructor() itinerary.Constructor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constructor")
	ret0, _ := ret[0].(itinerary.Constructor)
	return ret0
}

// Constructor indicates an expected call of Constructor.
func (mr *MockItineraryMockRecorder) Constructor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constructor", reflect.TypeOf((*MockItinerary)(nil).Constructor))
}

// Info mocks base method.
func (m *MockItinerary) Info() itinerary.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(itinerary.Info)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockItineraryMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockItinerary)(nil).Info))
}

// Modifier mocks base method.
func (m *MockItinerary) Modifier() itinerary.Modifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modifier")
	ret0, _ := ret[0].(itinerary.Modifier)
	return ret0
}

// Modifier indicates an expected call of Modifier.
func (mr *MockItineraryMockRecorder) Modifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modifier", reflect.TypeOf((*MockItinerary)(nil).Modifier))
}
