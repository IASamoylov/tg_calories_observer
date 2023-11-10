// Code generated by MockGen. DO NOT EDIT.
// Source: dependecies.go

// Package product is a generated GoMock package.
package product

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockinlineButton is a mock of inlineButton interface.
type MockinlineButton struct {
	ctrl     *gomock.Controller
	recorder *MockinlineButtonMockRecorder
}

// MockinlineButtonMockRecorder is the mock recorder for MockinlineButton.
type MockinlineButtonMockRecorder struct {
	mock *MockinlineButton
}

// NewMockinlineButton creates a new mock instance.
func NewMockinlineButton(ctrl *gomock.Controller) *MockinlineButton {
	mock := &MockinlineButton{ctrl: ctrl}
	mock.recorder = &MockinlineButtonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinlineButton) EXPECT() *MockinlineButtonMockRecorder {
	return m.recorder
}

// Callback mocks base method.
func (m *MockinlineButton) Callback() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Callback")
	ret0, _ := ret[0].(string)
	return ret0
}

// Callback indicates an expected call of Callback.
func (mr *MockinlineButtonMockRecorder) Callback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Callback", reflect.TypeOf((*MockinlineButton)(nil).Callback))
}

// name mocks base method.
func (m *MockinlineButton) Text() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "name")
	ret0, _ := ret[0].(string)
	return ret0
}

// name indicates an expected call of name.
func (mr *MockinlineButtonMockRecorder) Text() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "name", reflect.TypeOf((*MockinlineButton)(nil).Text))
}
