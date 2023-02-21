// Code generated by MockGen. DO NOT EDIT.
// Source: secrets-operator/internal/core/ports (interfaces: FindingsRepository,Notifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "secrets-operator/internal/core/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockFindingsRepository is a mock of FindingsRepository interface.
type MockFindingsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFindingsRepositoryMockRecorder
}

// MockFindingsRepositoryMockRecorder is the mock recorder for MockFindingsRepository.
type MockFindingsRepositoryMockRecorder struct {
	mock *MockFindingsRepository
}

// NewMockFindingsRepository creates a new mock instance.
func NewMockFindingsRepository(ctrl *gomock.Controller) *MockFindingsRepository {
	mock := &MockFindingsRepository{ctrl: ctrl}
	mock.recorder = &MockFindingsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindingsRepository) EXPECT() *MockFindingsRepositoryMockRecorder {
	return m.recorder
}

// GetRepoFindingsById mocks base method.
func (m *MockFindingsRepository) GetRepoFindingsById(arg0 int, arg1 string) (domain.RepoFindings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoFindingsById", arg0, arg1)
	ret0, _ := ret[0].(domain.RepoFindings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoFindingsById indicates an expected call of GetRepoFindingsById.
func (mr *MockFindingsRepositoryMockRecorder) GetRepoFindingsById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoFindingsById", reflect.TypeOf((*MockFindingsRepository)(nil).GetRepoFindingsById), arg0, arg1)
}

// GetRepositoriesByName mocks base method.
func (m *MockFindingsRepository) GetRepositoriesByName(arg0, arg1 string) ([]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoriesByName", arg0, arg1)
	ret0, _ := ret[0].([]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByName indicates an expected call of GetRepositoriesByName.
func (mr *MockFindingsRepositoryMockRecorder) GetRepositoriesByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoriesByName", reflect.TypeOf((*MockFindingsRepository)(nil).GetRepositoriesByName), arg0, arg1)
}

// SaveAndUpdateRepoFindingsById mocks base method.
func (m *MockFindingsRepository) SaveAndUpdateRepoFindingsById(arg0 domain.RepoFindings, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAndUpdateRepoFindingsById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAndUpdateRepoFindingsById indicates an expected call of SaveAndUpdateRepoFindingsById.
func (mr *MockFindingsRepositoryMockRecorder) SaveAndUpdateRepoFindingsById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAndUpdateRepoFindingsById", reflect.TypeOf((*MockFindingsRepository)(nil).SaveAndUpdateRepoFindingsById), arg0, arg1, arg2)
}

// SaveFindingsReport mocks base method.
func (m *MockFindingsRepository) SaveFindingsReport(arg0 domain.FindingsReport, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFindingsReport", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFindingsReport indicates an expected call of SaveFindingsReport.
func (mr *MockFindingsRepositoryMockRecorder) SaveFindingsReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFindingsReport", reflect.TypeOf((*MockFindingsRepository)(nil).SaveFindingsReport), arg0, arg1)
}

// MockNotifier is a mock of Notifier interface.
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier.
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance.
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockNotifier) SendMessage(arg0 domain.FindingsReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockNotifierMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockNotifier)(nil).SendMessage), arg0)
}
