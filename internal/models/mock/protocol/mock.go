// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/protocol/repository.go

// Package protocol is a generated GoMock package.
package protocol

import (
	model "github.com/baking-bad/bcdhub/internal/models/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(hash string, level int64) (model.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", hash, level)
	ret0, _ := ret[0].(model.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(hash, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), hash, level)
}

// GetAll mocks base method
func (m *MockRepository) GetAll() ([]model.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll))
}

// GetByNetworkWithSort mocks base method
func (m *MockRepository) GetByNetworkWithSort(sortField, order string) ([]model.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNetworkWithSort", sortField, order)
	ret0, _ := ret[0].([]model.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNetworkWithSort indicates an expected call of GetByNetworkWithSort
func (mr *MockRepositoryMockRecorder) GetByNetworkWithSort(sortField, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNetworkWithSort", reflect.TypeOf((*MockRepository)(nil).GetByNetworkWithSort), sortField, order)
}

// GetByID mocks base method
func (m *MockRepository) GetByID(id int64) (model.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), id)
}

// CheckChainID mocks base method
func (m *MockRepository) CheckChainID(chainID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChainID", chainID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckChainID indicates an expected call of CheckChainID
func (mr *MockRepositoryMockRecorder) CheckChainID(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChainID", reflect.TypeOf((*MockRepository)(nil).CheckChainID), chainID)
}
