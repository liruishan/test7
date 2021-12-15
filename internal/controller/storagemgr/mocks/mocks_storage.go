/*******************************************************************************
* Copyright 2021 Samsung Electronics All Rights Reserved.
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
// Source: storage.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return _m.recorder
}

// GetStatus mocks base method
func (_m *MockStorage) GetStatus() int {
	ret := _m.ctrl.Call(_m, "GetStatus")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetStatus indicates an expected call of GetStatus
func (_mr *MockStorageMockRecorder) GetStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetStatus", reflect.TypeOf((*MockStorage)(nil).GetStatus))
}

// StartStorage mocks base method
func (_m *MockStorage) StartStorage(host string) error {
	ret := _m.ctrl.Call(_m, "StartStorage", host)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartStorage indicates an expected call of StartStorage
func (_mr *MockStorageMockRecorder) StartStorage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "StartStorage", reflect.TypeOf((*MockStorage)(nil).StartStorage), arg0)
}

// BuildConfiguration mocks base method
func (_m *MockStorage) BuildConfiguration(host string) error {
	ret := _m.ctrl.Call(_m, "BuildConfiguration", host)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildConfiguration indicates an expected call of BuildConfiguration
func (_mr *MockStorageMockRecorder) BuildConfiguration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BuildConfiguration", reflect.TypeOf((*MockStorage)(nil).BuildConfiguration), arg0)
}
