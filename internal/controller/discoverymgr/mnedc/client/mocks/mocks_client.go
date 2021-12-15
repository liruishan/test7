/*******************************************************************************
 * Copyright 2020 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/lf-edge/edge-home-orchestration-go/internal/controller/discoverymgr/mnedc/client"
	client0 "github.com/lf-edge/edge-home-orchestration-go/internal/restinterface/client"
	reflect "reflect"
)

// MockMNEDCClient is a mock of MNEDCClient interface
type MockMNEDCClient struct {
	ctrl     *gomock.Controller
	recorder *MockMNEDCClientMockRecorder
}

// MockMNEDCClientMockRecorder is the mock recorder for MockMNEDCClient
type MockMNEDCClientMockRecorder struct {
	mock *MockMNEDCClient
}

// NewMockMNEDCClient creates a new mock instance
func NewMockMNEDCClient(ctrl *gomock.Controller) *MockMNEDCClient {
	mock := &MockMNEDCClient{ctrl: ctrl}
	mock.recorder = &MockMNEDCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMNEDCClient) EXPECT() *MockMNEDCClientMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockMNEDCClient) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockMNEDCClientMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMNEDCClient)(nil).Run))
}

// CreateClient mocks base method
func (m *MockMNEDCClient) CreateClient(arg0, arg1 string, arg2 bool) (*client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient
func (mr *MockMNEDCClientMockRecorder) CreateClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockMNEDCClient)(nil).CreateClient), arg0, arg1, arg2)
}

// Close mocks base method
func (m *MockMNEDCClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMNEDCClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMNEDCClient)(nil).Close))
}

// StartSendRoutine mocks base method
func (m *MockMNEDCClient) StartSendRoutine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartSendRoutine")
}

// StartSendRoutine indicates an expected call of StartSendRoutine
func (mr *MockMNEDCClientMockRecorder) StartSendRoutine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSendRoutine", reflect.TypeOf((*MockMNEDCClient)(nil).StartSendRoutine))
}

// StartRecvRoutine mocks base method
func (m *MockMNEDCClient) StartRecvRoutine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartRecvRoutine")
}

// StartRecvRoutine indicates an expected call of StartRecvRoutine
func (mr *MockMNEDCClientMockRecorder) StartRecvRoutine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRecvRoutine", reflect.TypeOf((*MockMNEDCClient)(nil).StartRecvRoutine))
}

// NotifyClose mocks base method
func (m *MockMNEDCClient) NotifyClose() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyClose")
}

// NotifyClose indicates an expected call of NotifyClose
func (mr *MockMNEDCClientMockRecorder) NotifyClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockMNEDCClient)(nil).NotifyClose))
}

// ConnectionReconciled mocks base method
func (m *MockMNEDCClient) ConnectionReconciled() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConnectionReconciled")
}

// ConnectionReconciled indicates an expected call of ConnectionReconciled
func (mr *MockMNEDCClientMockRecorder) ConnectionReconciled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionReconciled", reflect.TypeOf((*MockMNEDCClient)(nil).ConnectionReconciled))
}

// HandleError mocks base method
func (m *MockMNEDCClient) HandleError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleError", arg0)
}

// HandleError indicates an expected call of HandleError
func (mr *MockMNEDCClientMockRecorder) HandleError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleError", reflect.TypeOf((*MockMNEDCClient)(nil).HandleError), arg0)
}

// ParseVirtualIP mocks base method
func (m *MockMNEDCClient) ParseVirtualIP(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseVirtualIP", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseVirtualIP indicates an expected call of ParseVirtualIP
func (mr *MockMNEDCClientMockRecorder) ParseVirtualIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseVirtualIP", reflect.TypeOf((*MockMNEDCClient)(nil).ParseVirtualIP), arg0)
}

// TunReadRoutine mocks base method
func (m *MockMNEDCClient) TunReadRoutine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TunReadRoutine")
}

// TunReadRoutine indicates an expected call of TunReadRoutine
func (mr *MockMNEDCClientMockRecorder) TunReadRoutine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TunReadRoutine", reflect.TypeOf((*MockMNEDCClient)(nil).TunReadRoutine))
}

// TunWriteRoutine mocks base method
func (m *MockMNEDCClient) TunWriteRoutine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TunWriteRoutine")
}

// TunWriteRoutine indicates an expected call of TunWriteRoutine
func (mr *MockMNEDCClientMockRecorder) TunWriteRoutine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TunWriteRoutine", reflect.TypeOf((*MockMNEDCClient)(nil).TunWriteRoutine))
}

// NotifyBroadcastServer mocks base method
func (m *MockMNEDCClient) NotifyBroadcastServer(configPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyBroadcastServer", configPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyBroadcastServer indicates an expected call of NotifyBroadcastServer
func (mr *MockMNEDCClientMockRecorder) NotifyBroadcastServer(configPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyBroadcastServer", reflect.TypeOf((*MockMNEDCClient)(nil).NotifyBroadcastServer), configPath)
}

// SetClient mocks base method
func (m *MockMNEDCClient) SetClient(clientAPI client0.Clienter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", clientAPI)
}

// SetClient indicates an expected call of SetClient
func (mr *MockMNEDCClientMockRecorder) SetClient(clientAPI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockMNEDCClient)(nil).SetClient), clientAPI)
}
