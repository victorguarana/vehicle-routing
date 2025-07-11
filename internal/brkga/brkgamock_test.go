// Code generated by MockGen. DO NOT EDIT.
// Source: brkga.go
//
// Generated by this command:
//
//	mockgen -source=brkga.go -destination=brkgamock_test.go -package=brkga
//

// Package brkga is a generated GoMock package.
package brkga

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIDecoder is a mock of IDecoder interface.
type MockIDecoder[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIDecoderMockRecorder[T]
}

// MockIDecoderMockRecorder is the mock recorder for MockIDecoder.
type MockIDecoderMockRecorder[T any] struct {
	mock *MockIDecoder[T]
}

// NewMockIDecoder creates a new mock instance.
func NewMockIDecoder[T any](ctrl *gomock.Controller) *MockIDecoder[T] {
	mock := &MockIDecoder[T]{ctrl: ctrl}
	mock.recorder = &MockIDecoderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDecoder[T]) EXPECT() *MockIDecoderMockRecorder[T] {
	return m.recorder
}

// Decode mocks base method.
func (m *MockIDecoder[T]) Decode(arg0 *Individual) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockIDecoderMockRecorder[T]) Decode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockIDecoder[T])(nil).Decode), arg0)
}

// MockIMeasurer is a mock of IMeasurer interface.
type MockIMeasurer[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIMeasurerMockRecorder[T]
}

// MockIMeasurerMockRecorder is the mock recorder for MockIMeasurer.
type MockIMeasurerMockRecorder[T any] struct {
	mock *MockIMeasurer[T]
}

// NewMockIMeasurer creates a new mock instance.
func NewMockIMeasurer[T any](ctrl *gomock.Controller) *MockIMeasurer[T] {
	mock := &MockIMeasurer[T]{ctrl: ctrl}
	mock.recorder = &MockIMeasurerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMeasurer[T]) EXPECT() *MockIMeasurerMockRecorder[T] {
	return m.recorder
}

// Measure mocks base method.
func (m *MockIMeasurer[T]) Measure(arg0 T) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Measure", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Measure indicates an expected call of Measure.
func (mr *MockIMeasurerMockRecorder[T]) Measure(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Measure", reflect.TypeOf((*MockIMeasurer[T])(nil).Measure), arg0)
}
