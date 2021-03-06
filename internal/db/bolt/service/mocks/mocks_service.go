/*******************************************************************************
 * Copyright 2019 Samsung Electronics All Rights Reserved.
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
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	service "github.com/lf-edge/edge-home-orchestration-go/internal/db/bolt/service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDBInterface is a mock of DBInterface interface
type MockDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBInterfaceMockRecorder
}

// MockDBInterfaceMockRecorder is the mock recorder for MockDBInterface
type MockDBInterfaceMockRecorder struct {
	mock *MockDBInterface
}

// NewMockDBInterface creates a new mock instance
func NewMockDBInterface(ctrl *gomock.Controller) *MockDBInterface {
	mock := &MockDBInterface{ctrl: ctrl}
	mock.recorder = &MockDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBInterface) EXPECT() *MockDBInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDBInterface) Get(id string) (service.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(service.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDBInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBInterface)(nil).Get), id)
}

// GetList mocks base method
func (m *MockDBInterface) GetList() ([]service.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]service.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList
func (mr *MockDBInterfaceMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockDBInterface)(nil).GetList))
}

// Set mocks base method
func (m *MockDBInterface) Set(info service.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockDBInterfaceMockRecorder) Set(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDBInterface)(nil).Set), info)
}

// Update mocks base method
func (m *MockDBInterface) Update(info service.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDBInterfaceMockRecorder) Update(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDBInterface)(nil).Update), info)
}

// Delete mocks base method
func (m *MockDBInterface) Delete(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDBInterfaceMockRecorder) Delete(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDBInterface)(nil).Delete), name)
}
