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
// Source: notification.go

// Package mock_notification is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/lf-edge/edge-home-orchestration-go/internal/restinterface/client"
	reflect "reflect"
)

// MockNotification is a mock of Notification interface
type MockNotification struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationMockRecorder
}

// MockNotificationMockRecorder is the mock recorder for MockNotification
type MockNotificationMockRecorder struct {
	mock *MockNotification
}

// NewMockNotification creates a new mock instance
func NewMockNotification(ctrl *gomock.Controller) *MockNotification {
	mock := &MockNotification{ctrl: ctrl}
	mock.recorder = &MockNotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotification) EXPECT() *MockNotificationMockRecorder {
	return m.recorder
}

// InvokeNotification mocks base method
func (m *MockNotification) InvokeNotification(target string, serviceID float64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeNotification", target, serviceID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeNotification indicates an expected call of InvokeNotification
func (mr *MockNotificationMockRecorder) InvokeNotification(target, serviceID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeNotification", reflect.TypeOf((*MockNotification)(nil).InvokeNotification), target, serviceID, status)
}

// AddNotificationChan mocks base method
func (m *MockNotification) AddNotificationChan(serviceID uint64, notiChan chan string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddNotificationChan", serviceID, notiChan)
}

// AddNotificationChan indicates an expected call of AddNotificationChan
func (mr *MockNotificationMockRecorder) AddNotificationChan(serviceID, notiChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNotificationChan", reflect.TypeOf((*MockNotification)(nil).AddNotificationChan), serviceID, notiChan)
}

// HandleNotificationOnLocal mocks base method
func (m *MockNotification) HandleNotificationOnLocal(serviceID float64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNotificationOnLocal", serviceID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleNotificationOnLocal indicates an expected call of HandleNotificationOnLocal
func (mr *MockNotificationMockRecorder) HandleNotificationOnLocal(serviceID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNotificationOnLocal", reflect.TypeOf((*MockNotification)(nil).HandleNotificationOnLocal), serviceID, status)
}

// SetClient mocks base method
func (m *MockNotification) SetClient(clientAPI client.Clienter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", clientAPI)
}

// SetClient indicates an expected call of SetClient
func (mr *MockNotificationMockRecorder) SetClient(clientAPI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockNotification)(nil).SetClient), clientAPI)
}
