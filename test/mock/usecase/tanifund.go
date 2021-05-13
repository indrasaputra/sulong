// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/tanifund.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/indrasaputra/sulong/entity"
)

// MockTaniFundProjectGetter is a mock of TaniFundProjectGetter interface
type MockTaniFundProjectGetter struct {
	ctrl     *gomock.Controller
	recorder *MockTaniFundProjectGetterMockRecorder
}

// MockTaniFundProjectGetterMockRecorder is the mock recorder for MockTaniFundProjectGetter
type MockTaniFundProjectGetterMockRecorder struct {
	mock *MockTaniFundProjectGetter
}

// NewMockTaniFundProjectGetter creates a new mock instance
func NewMockTaniFundProjectGetter(ctrl *gomock.Controller) *MockTaniFundProjectGetter {
	mock := &MockTaniFundProjectGetter{ctrl: ctrl}
	mock.recorder = &MockTaniFundProjectGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaniFundProjectGetter) EXPECT() *MockTaniFundProjectGetterMockRecorder {
	return m.recorder
}

// GetNewestProjects mocks base method
func (m *MockTaniFundProjectGetter) GetNewestProjects(ctx context.Context, numberOfProject int) ([]*entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewestProjects", ctx, numberOfProject)
	ret0, _ := ret[0].([]*entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewestProjects indicates an expected call of GetNewestProjects
func (mr *MockTaniFundProjectGetterMockRecorder) GetNewestProjects(ctx, numberOfProject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewestProjects", reflect.TypeOf((*MockTaniFundProjectGetter)(nil).GetNewestProjects), ctx, numberOfProject)
}

// MockTaniFundProjectNotifier is a mock of TaniFundProjectNotifier interface
type MockTaniFundProjectNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockTaniFundProjectNotifierMockRecorder
}

// MockTaniFundProjectNotifierMockRecorder is the mock recorder for MockTaniFundProjectNotifier
type MockTaniFundProjectNotifierMockRecorder struct {
	mock *MockTaniFundProjectNotifier
}

// NewMockTaniFundProjectNotifier creates a new mock instance
func NewMockTaniFundProjectNotifier(ctrl *gomock.Controller) *MockTaniFundProjectNotifier {
	mock := &MockTaniFundProjectNotifier{ctrl: ctrl}
	mock.recorder = &MockTaniFundProjectNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaniFundProjectNotifier) EXPECT() *MockTaniFundProjectNotifierMockRecorder {
	return m.recorder
}

// Notify mocks base method
func (m *MockTaniFundProjectNotifier) Notify(ctx context.Context, project *entity.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify
func (mr *MockTaniFundProjectNotifierMockRecorder) Notify(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockTaniFundProjectNotifier)(nil).Notify), ctx, project)
}
