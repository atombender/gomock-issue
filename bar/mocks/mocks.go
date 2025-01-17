// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/atombender/gomock-issue/bar (interfaces: Starter)
//
// Generated by this command:
//
//	mockgen -package mocks -destination ./bar/mocks/mocks.go github.com/atombender/gomock-issue/bar Starter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	things "github.com/atombender/gomock-issue/foo/internal/things"
	gomock "go.uber.org/mock/gomock"
)

// MockStarter is a mock of Starter interface.
type MockStarter struct {
	ctrl     *gomock.Controller
	recorder *MockStarterMockRecorder
}

// MockStarterMockRecorder is the mock recorder for MockStarter.
type MockStarterMockRecorder struct {
	mock *MockStarter
}

// NewMockStarter creates a new mock instance.
func NewMockStarter(ctrl *gomock.Controller) *MockStarter {
	mock := &MockStarter{ctrl: ctrl}
	mock.recorder = &MockStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStarter) EXPECT() *MockStarterMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockStarter) Start(arg0 *things.Thing) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockStarterMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStarter)(nil).Start), arg0)
}
