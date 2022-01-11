// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaimu/speedtest/providers/netflix (interfaces: Netflix)

// Package mock_netflix is a generated GoMock package.
package mock_netflix

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetflix is a mock of Netflix interface.
type MockNetflix struct {
	ctrl     *gomock.Controller
	recorder *MockNetflixMockRecorder
}

// MockNetflixMockRecorder is the mock recorder for MockNetflix.
type MockNetflixMockRecorder struct {
	mock *MockNetflix
}

// NewMockNetflix creates a new mock instance.
func NewMockNetflix(ctrl *gomock.Controller) *MockNetflix {
	mock := &MockNetflix{ctrl: ctrl}
	mock.recorder = &MockNetflixMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetflix) EXPECT() *MockNetflixMockRecorder {
	return m.recorder
}

// GetUrls mocks base method.
func (m *MockNetflix) GetUrls() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrls")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrls indicates an expected call of GetUrls.
func (mr *MockNetflixMockRecorder) GetUrls() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrls", reflect.TypeOf((*MockNetflix)(nil).GetUrls))
}

// Init mocks base method.
func (m *MockNetflix) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockNetflixMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockNetflix)(nil).Init))
}

// Measure mocks base method.
func (m *MockNetflix) Measure(arg0 []string, arg1 chan<- float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Measure", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Measure indicates an expected call of Measure.
func (mr *MockNetflixMockRecorder) Measure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Measure", reflect.TypeOf((*MockNetflix)(nil).Measure), arg0, arg1)
}
