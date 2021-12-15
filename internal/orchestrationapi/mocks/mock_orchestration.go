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
// Source: orchestrationapi (interfaces: Orche,OrcheExternalAPI,OrcheInternalAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	configuremgrtypes "github.com/lf-edge/edge-home-orchestration-go/internal/common/types/configuremgrtypes"
	verifier "github.com/lf-edge/edge-home-orchestration-go/internal/controller/securemgr/verifier"
	orchestrationapi "github.com/lf-edge/edge-home-orchestration-go/internal/orchestrationapi"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrche is a mock of Orche interface.
type MockOrche struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheMockRecorder
}

// MockOrcheMockRecorder is the mock recorder for MockOrche.
type MockOrcheMockRecorder struct {
	mock *MockOrche
}

// NewMockOrche creates a new mock instance.
func NewMockOrche(ctrl *gomock.Controller) *MockOrche {
	mock := &MockOrche{ctrl: ctrl}
	mock.recorder = &MockOrcheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrche) EXPECT() *MockOrcheMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockOrche) Start(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0, arg1, arg2)
}

// Start indicates an expected call of Start.
func (mr *MockOrcheMockRecorder) Start(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOrche)(nil).Start), arg0, arg1, arg2)
}

// MockOrcheExternalAPI is a mock of OrcheExternalAPI interface.
type MockOrcheExternalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheExternalAPIMockRecorder
}

// MockOrcheExternalAPIMockRecorder is the mock recorder for MockOrcheExternalAPI.
type MockOrcheExternalAPIMockRecorder struct {
	mock *MockOrcheExternalAPI
}

// NewMockOrcheExternalAPI creates a new mock instance.
func NewMockOrcheExternalAPI(ctrl *gomock.Controller) *MockOrcheExternalAPI {
	mock := &MockOrcheExternalAPI{ctrl: ctrl}
	mock.recorder = &MockOrcheExternalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrcheExternalAPI) EXPECT() *MockOrcheExternalAPIMockRecorder {
	return m.recorder
}

// RequestService mocks base method.
func (m *MockOrcheExternalAPI) RequestService(arg0 orchestrationapi.ReqeustService) orchestrationapi.ResponseService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestService", arg0)
	ret0, _ := ret[0].(orchestrationapi.ResponseService)
	return ret0
}

// RequestService indicates an expected call of RequestService.
func (mr *MockOrcheExternalAPIMockRecorder) RequestService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestService", reflect.TypeOf((*MockOrcheExternalAPI)(nil).RequestService), arg0)
}

// RequestVerifierConf mocks base method.
func (m *MockOrcheExternalAPI) RequestVerifierConf(arg0 verifier.RequestVerifierConf) verifier.ResponseVerifierConf {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestVerifierConf", arg0)
	ret0, _ := ret[0].(verifier.ResponseVerifierConf)
	return ret0
}

// RequestVerifierConf indicates an expected call of RequestVerifierConf.
func (mr *MockOrcheExternalAPIMockRecorder) RequestVerifierConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestVerifierConf", reflect.TypeOf((*MockOrcheExternalAPI)(nil).RequestVerifierConf), arg0)
}

// MockOrcheInternalAPI is a mock of OrcheInternalAPI interface.
type MockOrcheInternalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheInternalAPIMockRecorder
}

// MockOrcheInternalAPIMockRecorder is the mock recorder for MockOrcheInternalAPI.
type MockOrcheInternalAPIMockRecorder struct {
	mock *MockOrcheInternalAPI
}

// NewMockOrcheInternalAPI creates a new mock instance.
func NewMockOrcheInternalAPI(ctrl *gomock.Controller) *MockOrcheInternalAPI {
	mock := &MockOrcheInternalAPI{ctrl: ctrl}
	mock.recorder = &MockOrcheInternalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrcheInternalAPI) EXPECT() *MockOrcheInternalAPIMockRecorder {
	return m.recorder
}

// ExecuteAppOnLocal mocks base method.
func (m *MockOrcheInternalAPI) ExecuteAppOnLocal(arg0 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAppOnLocal", arg0)
}

// ExecuteAppOnLocal indicates an expected call of ExecuteAppOnLocal.
func (mr *MockOrcheInternalAPIMockRecorder) ExecuteAppOnLocal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAppOnLocal", reflect.TypeOf((*MockOrcheInternalAPI)(nil).ExecuteAppOnLocal), arg0)
}

// GetOrchestrationInfo mocks base method.
func (m *MockOrcheInternalAPI) GetOrchestrationInfo() (string, string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrchestrationInfo")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetOrchestrationInfo indicates an expected call of GetOrchestrationInfo.
func (mr *MockOrcheInternalAPIMockRecorder) GetOrchestrationInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrchestrationInfo", reflect.TypeOf((*MockOrcheInternalAPI)(nil).GetOrchestrationInfo))
}

// GetResource mocks base method.
func (m *MockOrcheInternalAPI) GetResource(arg0 string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockOrcheInternalAPIMockRecorder) GetResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockOrcheInternalAPI)(nil).GetResource), arg0)
}

// GetScore mocks base method.
func (m *MockOrcheInternalAPI) GetScore(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScore", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScore indicates an expected call of GetScore.
func (mr *MockOrcheInternalAPIMockRecorder) GetScore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScore", reflect.TypeOf((*MockOrcheInternalAPI)(nil).GetScore), arg0)
}

// GetScoreWithResource mocks base method.
func (m *MockOrcheInternalAPI) GetScoreWithResource(arg0 map[string]interface{}) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScoreWithResource", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScoreWithResource indicates an expected call of GetScoreWithResource.
func (mr *MockOrcheInternalAPIMockRecorder) GetScoreWithResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScoreWithResource", reflect.TypeOf((*MockOrcheInternalAPI)(nil).GetScoreWithResource), arg0)
}

// HandleDeviceInfo mocks base method.
func (m *MockOrcheInternalAPI) HandleDeviceInfo(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleDeviceInfo", arg0, arg1, arg2)
}

// HandleDeviceInfo indicates an expected call of HandleDeviceInfo.
func (mr *MockOrcheInternalAPIMockRecorder) HandleDeviceInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeviceInfo", reflect.TypeOf((*MockOrcheInternalAPI)(nil).HandleDeviceInfo), arg0, arg1, arg2)
}

// HandleNotificationOnLocal mocks base method.
func (m *MockOrcheInternalAPI) HandleNotificationOnLocal(arg0 float64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNotificationOnLocal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleNotificationOnLocal indicates an expected call of HandleNotificationOnLocal.
func (mr *MockOrcheInternalAPIMockRecorder) HandleNotificationOnLocal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNotificationOnLocal", reflect.TypeOf((*MockOrcheInternalAPI)(nil).HandleNotificationOnLocal), arg0, arg1)
}

// Notify mocks base method.
func (m *MockOrcheInternalAPI) Notify(arg0 configuremgrtypes.ServiceInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", arg0)
}

// Notify indicates an expected call of Notify.
func (mr *MockOrcheInternalAPIMockRecorder) Notify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockOrcheInternalAPI)(nil).Notify), arg0)
}
