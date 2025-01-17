// Code generated by MockGen. DO NOT EDIT.
// Source: internal/noderpc/interface.go

// Package noderpc is a generated GoMock package.
package noderpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockINode is a mock of INode interface
type MockINode struct {
	ctrl     *gomock.Controller
	recorder *MockINodeMockRecorder
}

// MockINodeMockRecorder is the mock recorder for MockINode
type MockINodeMockRecorder struct {
	mock *MockINode
}

// NewMockINode creates a new mock instance
func NewMockINode(ctrl *gomock.Controller) *MockINode {
	mock := &MockINode{ctrl: ctrl}
	mock.recorder = &MockINodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINode) EXPECT() *MockINodeMockRecorder {
	return m.recorder
}

// Block mocks base method
func (m *MockINode) Block(arg0 context.Context, arg1 int64) (Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", arg0, arg1)
	ret0, _ := ret[0].(Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block
func (mr *MockINodeMockRecorder) Block(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockINode)(nil).Block), arg0, arg1)
}

// GetHead mocks base method
func (m *MockINode) GetHead(arg0 context.Context) (Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", arg0)
	ret0, _ := ret[0].(Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead
func (mr *MockINodeMockRecorder) GetHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockINode)(nil).GetHead), arg0)
}

// GetHeader mocks base method
func (m *MockINode) GetHeader(arg0 context.Context, arg1 int64) (Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockINodeMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockINode)(nil).GetHeader), arg0, arg1)
}

// GetScriptJSON mocks base method
func (m *MockINode) GetScriptJSON(arg0 context.Context, arg1 string, arg2 int64) (Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScriptJSON", arg0, arg1, arg2)
	ret0, _ := ret[0].(Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptJSON indicates an expected call of GetScriptJSON
func (mr *MockINodeMockRecorder) GetScriptJSON(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptJSON", reflect.TypeOf((*MockINode)(nil).GetScriptJSON), arg0, arg1, arg2)
}

// GetRawScript mocks base method
func (m *MockINode) GetRawScript(ctx context.Context, address string, level int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawScript", ctx, address, level)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawScript indicates an expected call of GetRawScript
func (mr *MockINodeMockRecorder) GetRawScript(ctx, address, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawScript", reflect.TypeOf((*MockINode)(nil).GetRawScript), ctx, address, level)
}

// GetScriptStorageRaw mocks base method
func (m *MockINode) GetScriptStorageRaw(arg0 context.Context, arg1 string, arg2 int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScriptStorageRaw", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptStorageRaw indicates an expected call of GetScriptStorageRaw
func (mr *MockINodeMockRecorder) GetScriptStorageRaw(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptStorageRaw", reflect.TypeOf((*MockINode)(nil).GetScriptStorageRaw), arg0, arg1, arg2)
}

// GetContractBalance mocks base method
func (m *MockINode) GetContractBalance(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractBalance indicates an expected call of GetContractBalance
func (mr *MockINodeMockRecorder) GetContractBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractBalance", reflect.TypeOf((*MockINode)(nil).GetContractBalance), arg0, arg1, arg2)
}

// GetContractData mocks base method
func (m *MockINode) GetContractData(arg0 context.Context, arg1 string, arg2 int64) (ContractData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractData", arg0, arg1, arg2)
	ret0, _ := ret[0].(ContractData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractData indicates an expected call of GetContractData
func (mr *MockINodeMockRecorder) GetContractData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractData", reflect.TypeOf((*MockINode)(nil).GetContractData), arg0, arg1, arg2)
}

// GetOPG mocks base method
func (m *MockINode) GetOPG(ctx context.Context, block int64) ([]OperationGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOPG", ctx, block)
	ret0, _ := ret[0].([]OperationGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOPG indicates an expected call of GetOPG
func (mr *MockINodeMockRecorder) GetOPG(ctx, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOPG", reflect.TypeOf((*MockINode)(nil).GetOPG), ctx, block)
}

// GetLightOPG mocks base method
func (m *MockINode) GetLightOPG(ctx context.Context, block int64) ([]LightOperationGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLightOPG", ctx, block)
	ret0, _ := ret[0].([]LightOperationGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLightOPG indicates an expected call of GetLightOPG
func (mr *MockINodeMockRecorder) GetLightOPG(ctx, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLightOPG", reflect.TypeOf((*MockINode)(nil).GetLightOPG), ctx, block)
}

// GetContractsByBlock mocks base method
func (m *MockINode) GetContractsByBlock(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractsByBlock", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractsByBlock indicates an expected call of GetContractsByBlock
func (mr *MockINodeMockRecorder) GetContractsByBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractsByBlock", reflect.TypeOf((*MockINode)(nil).GetContractsByBlock), arg0, arg1)
}

// GetNetworkConstants mocks base method
func (m *MockINode) GetNetworkConstants(arg0 context.Context, arg1 int64) (Constants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkConstants", arg0, arg1)
	ret0, _ := ret[0].(Constants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkConstants indicates an expected call of GetNetworkConstants
func (mr *MockINodeMockRecorder) GetNetworkConstants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkConstants", reflect.TypeOf((*MockINode)(nil).GetNetworkConstants), arg0, arg1)
}

// RunCode mocks base method
func (m *MockINode) RunCode(arg0 context.Context, arg1, arg2, arg3 []byte, arg4, arg5, arg6, arg7, arg8 string, arg9, arg10 int64) (RunCodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCode", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	ret0, _ := ret[0].(RunCodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCode indicates an expected call of RunCode
func (mr *MockINodeMockRecorder) RunCode(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCode", reflect.TypeOf((*MockINode)(nil).RunCode), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RunOperation mocks base method
func (m *MockINode) RunOperation(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5, arg6, arg7, arg8, arg9 int64, arg10 []byte) (OperationGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOperation", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	ret0, _ := ret[0].(OperationGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOperation indicates an expected call of RunOperation
func (mr *MockINodeMockRecorder) RunOperation(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOperation", reflect.TypeOf((*MockINode)(nil).RunOperation), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RunOperationLight mocks base method
func (m *MockINode) RunOperationLight(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5, arg6, arg7, arg8, arg9 int64, arg10 []byte) (LightOperationGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOperationLight", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	ret0, _ := ret[0].(LightOperationGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOperationLight indicates an expected call of RunOperationLight
func (mr *MockINodeMockRecorder) RunOperationLight(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOperationLight", reflect.TypeOf((*MockINode)(nil).RunOperationLight), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RunScriptView mocks base method
func (m *MockINode) RunScriptView(ctx context.Context, request RunScriptViewRequest) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunScriptView", ctx, request)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunScriptView indicates an expected call of RunScriptView
func (mr *MockINodeMockRecorder) RunScriptView(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunScriptView", reflect.TypeOf((*MockINode)(nil).RunScriptView), ctx, request)
}

// GetCounter mocks base method
func (m *MockINode) GetCounter(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounter indicates an expected call of GetCounter
func (mr *MockINodeMockRecorder) GetCounter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockINode)(nil).GetCounter), arg0, arg1)
}

// GetBigMapType mocks base method
func (m *MockINode) GetBigMapType(ctx context.Context, ptr, level int64) (BigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBigMapType", ctx, ptr, level)
	ret0, _ := ret[0].(BigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBigMapType indicates an expected call of GetBigMapType
func (mr *MockINodeMockRecorder) GetBigMapType(ctx, ptr, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBigMapType", reflect.TypeOf((*MockINode)(nil).GetBigMapType), ctx, ptr, level)
}

// GetBlockMetadata mocks base method
func (m *MockINode) GetBlockMetadata(ctx context.Context, level int64) (Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockMetadata", ctx, level)
	ret0, _ := ret[0].(Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockMetadata indicates an expected call of GetBlockMetadata
func (mr *MockINodeMockRecorder) GetBlockMetadata(ctx, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockMetadata", reflect.TypeOf((*MockINode)(nil).GetBlockMetadata), ctx, level)
}
