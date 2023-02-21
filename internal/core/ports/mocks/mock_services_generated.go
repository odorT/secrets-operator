// Code generated by MockGen. DO NOT EDIT.
// Source: secrets-operator/internal/core/ports (interfaces: FindingService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "secrets-operator/internal/core/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockFindingService is a mock of FindingService interface.
type MockFindingService struct {
	ctrl     *gomock.Controller
	recorder *MockFindingServiceMockRecorder
}

// MockFindingServiceMockRecorder is the mock recorder for MockFindingService.
type MockFindingServiceMockRecorder struct {
	mock *MockFindingService
}

// NewMockFindingService creates a new mock instance.
func NewMockFindingService(ctrl *gomock.Controller) *MockFindingService {
	mock := &MockFindingService{ctrl: ctrl}
	mock.recorder = &MockFindingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindingService) EXPECT() *MockFindingServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockFindingService) Add(arg0 domain.FindingsReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockFindingServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFindingService)(nil).Add), arg0)
}

// GetById mocks base method.
func (m *MockFindingService) GetById(arg0 int) (domain.RepoFindings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(domain.RepoFindings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockFindingServiceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockFindingService)(nil).GetById), arg0)
}

// GetByName mocks base method.
func (m *MockFindingService) GetByName(arg0 string) ([]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0)
	ret0, _ := ret[0].([]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockFindingServiceMockRecorder) GetByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockFindingService)(nil).GetByName), arg0)
}

// Notify mocks base method.
func (m *MockFindingService) Notify(arg0 domain.FindingsReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockFindingServiceMockRecorder) Notify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockFindingService)(nil).Notify), arg0)
}
