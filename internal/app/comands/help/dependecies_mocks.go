// Code generated by MockGen. DO NOT EDIT.
// Source: dependecies.go

// Package help is a generated GoMock package.
package help

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockсommand is a mock of сommand interface.
type Mockсommand struct {
	ctrl     *gomock.Controller
	recorder *MockсommandMockRecorder
}

// MockсommandMockRecorder is the mock recorder for Mockсommand.
type MockсommandMockRecorder struct {
	mock *Mockсommand
}

// NewMockсommand creates a new mock instance.
func NewMockсommand(ctrl *gomock.Controller) *Mockсommand {
	mock := &Mockсommand{ctrl: ctrl}
	mock.recorder = &MockсommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockсommand) EXPECT() *MockсommandMockRecorder {
	return m.recorder
}

// Alias mocks base method.
func (m *Mockсommand) Alias() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alias")
	ret0, _ := ret[0].(string)
	return ret0
}

// Alias indicates an expected call of Alias.
func (mr *MockсommandMockRecorder) Alias() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alias", reflect.TypeOf((*Mockсommand)(nil).Alias))
}

// Description mocks base method.
func (m *Mockсommand) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockсommandMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*Mockсommand)(nil).Description))
}
