// Code generated by MockGen. DO NOT EDIT.
// Source: timer.go

// Package mock_timerwatcher is a generated GoMock package.
package timerwatcher

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// State mocks base method
func (m *MockMachine) State() string {
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(string)
	return ret0
}

// State indicates an expected call of State
func (mr *MockMachineMockRecorder) State() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockMachine)(nil).State))
}

// Name mocks base method
func (m *MockMachine) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockMachineMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMachine)(nil).Name))
}

// Identity mocks base method
func (m *MockMachine) Identity() string {
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity
func (mr *MockMachineMockRecorder) Identity() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockMachine)(nil).Identity))
}

// InstanceID mocks base method
func (m *MockMachine) InstanceID() string {
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID
func (mr *MockMachineMockRecorder) InstanceID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachine)(nil).InstanceID))
}

// Version mocks base method
func (m *MockMachine) Version() string {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockMachineMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMachine)(nil).Version))
}

// TimeStampSeconds mocks base method
func (m *MockMachine) TimeStampSeconds() int64 {
	ret := m.ctrl.Call(m, "TimeStampSeconds")
	ret0, _ := ret[0].(int64)
	return ret0
}

// TimeStampSeconds indicates an expected call of TimeStampSeconds
func (mr *MockMachineMockRecorder) TimeStampSeconds() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeStampSeconds", reflect.TypeOf((*MockMachine)(nil).TimeStampSeconds))
}

// NotifyWatcherState mocks base method
func (m *MockMachine) NotifyWatcherState(arg0 string, arg1 interface{}) {
	m.ctrl.Call(m, "NotifyWatcherState", arg0, arg1)
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState
func (mr *MockMachineMockRecorder) NotifyWatcherState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockMachine)(nil).NotifyWatcherState), arg0, arg1)
}

// Transition mocks base method
func (m *MockMachine) Transition(t string, args ...interface{}) error {
	varargs := []interface{}{t}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transition indicates an expected call of Transition
func (mr *MockMachineMockRecorder) Transition(t interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{t}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transition", reflect.TypeOf((*MockMachine)(nil).Transition), varargs...)
}

// Debugf mocks base method
func (m *MockMachine) Debugf(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockMachineMockRecorder) Debugf(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockMachine)(nil).Debugf), varargs...)
}

// Infof mocks base method
func (m *MockMachine) Infof(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockMachineMockRecorder) Infof(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockMachine)(nil).Infof), varargs...)
}

// Errorf mocks base method
func (m *MockMachine) Errorf(name, format string, args ...interface{}) {
	varargs := []interface{}{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockMachineMockRecorder) Errorf(name, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockMachine)(nil).Errorf), varargs...)
}
