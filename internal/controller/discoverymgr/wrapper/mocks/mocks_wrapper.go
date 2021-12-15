// Code generated by MockGen. DO NOT EDIT.
// Source: controller/discoverymgr/wrapper (interfaces: ZeroconfInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	wrapper "github.com/lf-edge/edge-home-orchestration-go/internal/controller/discoverymgr/wrapper"
	net "net"
	reflect "reflect"
)

// MockZeroconfInterface is a mock of ZeroconfInterface interface
type MockZeroconfInterface struct {
	ctrl     *gomock.Controller
	recorder *MockZeroconfInterfaceMockRecorder
}

// MockZeroconfInterfaceMockRecorder is the mock recorder for MockZeroconfInterface
type MockZeroconfInterfaceMockRecorder struct {
	mock *MockZeroconfInterface
}

// NewMockZeroconfInterface creates a new mock instance
func NewMockZeroconfInterface(ctrl *gomock.Controller) *MockZeroconfInterface {
	mock := &MockZeroconfInterface{ctrl: ctrl}
	mock.recorder = &MockZeroconfInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockZeroconfInterface) EXPECT() *MockZeroconfInterfaceMockRecorder {
	return m.recorder
}

// Advertise mocks base method
func (m *MockZeroconfInterface) Advertise() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Advertise")
}

// Advertise indicates an expected call of Advertise
func (mr *MockZeroconfInterfaceMockRecorder) Advertise() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Advertise", reflect.TypeOf((*MockZeroconfInterface)(nil).Advertise))
}

// GetSubscriberChan mocks base method
func (m *MockZeroconfInterface) GetSubscriberChan() (chan *wrapper.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriberChan")
	ret0, _ := ret[0].(chan *wrapper.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriberChan indicates an expected call of GetSubscriberChan
func (mr *MockZeroconfInterfaceMockRecorder) GetSubscriberChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriberChan", reflect.TypeOf((*MockZeroconfInterface)(nil).GetSubscriberChan))
}

// GetText mocks base method
func (m *MockZeroconfInterface) GetText() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetText")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetText indicates an expected call of GetText
func (mr *MockZeroconfInterfaceMockRecorder) GetText() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetText", reflect.TypeOf((*MockZeroconfInterface)(nil).GetText))
}

// RegisterProxy mocks base method
func (m *MockZeroconfInterface) RegisterProxy(arg0, arg1, arg2 string, arg3 int, arg4 string, arg5, arg6 []string, arg7 []net.Interface) (wrapper.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterProxy", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(wrapper.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterProxy indicates an expected call of RegisterProxy
func (mr *MockZeroconfInterfaceMockRecorder) RegisterProxy(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterProxy", reflect.TypeOf((*MockZeroconfInterface)(nil).RegisterProxy), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ResetServer mocks base method
func (m *MockZeroconfInterface) ResetServer(arg0 []net.IP) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetServer", arg0)
}

// ResetServer indicates an expected call of ResetServer
func (mr *MockZeroconfInterfaceMockRecorder) ResetServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetServer", reflect.TypeOf((*MockZeroconfInterface)(nil).ResetServer), arg0)
}

// SetText mocks base method
func (m *MockZeroconfInterface) SetText(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetText", arg0)
}

// SetText indicates an expected call of SetText
func (mr *MockZeroconfInterfaceMockRecorder) SetText(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetText", reflect.TypeOf((*MockZeroconfInterface)(nil).SetText), arg0)
}

// Shutdown mocks base method
func (m *MockZeroconfInterface) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockZeroconfInterfaceMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockZeroconfInterface)(nil).Shutdown))
}
