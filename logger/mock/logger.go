// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ashokhirpara1/golang-prettylogs/logger (interfaces: Storage)

// Package mocklogger is a generated GoMock package.
package mocklogger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockStorage) Error(arg0 logrus.Fields, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0, arg1)
}

// Error indicates an expected call of Error.
func (mr *MockStorageMockRecorder) Error(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockStorage)(nil).Error), arg0, arg1)
}

// Fatal mocks base method.
func (m *MockStorage) Fatal(arg0 logrus.Fields, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fatal", arg0, arg1)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockStorageMockRecorder) Fatal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockStorage)(nil).Fatal), arg0, arg1)
}

// Info mocks base method.
func (m *MockStorage) Info(arg0 logrus.Fields, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", arg0, arg1)
}

// Info indicates an expected call of Info.
func (mr *MockStorageMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockStorage)(nil).Info), arg0, arg1)
}
