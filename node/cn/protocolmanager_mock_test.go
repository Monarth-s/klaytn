// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Monarth-s/klaytn/node/cn (interfaces: BackendProtocolManager)

// Package cn is a generated GoMock package.
package cn

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	accounts "github.com/Monarth-s/klaytn/accounts"
	types "github.com/Monarth-s/klaytn/blockchain/types"
	common "github.com/Monarth-s/klaytn/common"
	p2p "github.com/Monarth-s/klaytn/networks/p2p"
)

// MockBackendProtocolManager is a mock of BackendProtocolManager interface.
type MockBackendProtocolManager struct {
	ctrl     *gomock.Controller
	recorder *MockBackendProtocolManagerMockRecorder
}

// MockBackendProtocolManagerMockRecorder is the mock recorder for MockBackendProtocolManager.
type MockBackendProtocolManagerMockRecorder struct {
	mock *MockBackendProtocolManager
}

// NewMockBackendProtocolManager creates a new mock instance.
func NewMockBackendProtocolManager(ctrl *gomock.Controller) *MockBackendProtocolManager {
	mock := &MockBackendProtocolManager{ctrl: ctrl}
	mock.recorder = &MockBackendProtocolManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendProtocolManager) EXPECT() *MockBackendProtocolManagerMockRecorder {
	return m.recorder
}

// Downloader mocks base method.
func (m *MockBackendProtocolManager) Downloader() ProtocolManagerDownloader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Downloader")
	ret0, _ := ret[0].(ProtocolManagerDownloader)
	return ret0
}

// Downloader indicates an expected call of Downloader.
func (mr *MockBackendProtocolManagerMockRecorder) Downloader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Downloader", reflect.TypeOf((*MockBackendProtocolManager)(nil).Downloader))
}

// GetSubProtocols mocks base method.
func (m *MockBackendProtocolManager) GetSubProtocols() []p2p.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubProtocols")
	ret0, _ := ret[0].([]p2p.Protocol)
	return ret0
}

// GetSubProtocols indicates an expected call of GetSubProtocols.
func (mr *MockBackendProtocolManagerMockRecorder) GetSubProtocols() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubProtocols", reflect.TypeOf((*MockBackendProtocolManager)(nil).GetSubProtocols))
}

// NodeType mocks base method.
func (m *MockBackendProtocolManager) NodeType() common.ConnType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeType")
	ret0, _ := ret[0].(common.ConnType)
	return ret0
}

// NodeType indicates an expected call of NodeType.
func (mr *MockBackendProtocolManagerMockRecorder) NodeType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeType", reflect.TypeOf((*MockBackendProtocolManager)(nil).NodeType))
}

// ProtocolVersion mocks base method.
func (m *MockBackendProtocolManager) ProtocolVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtocolVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// ProtocolVersion indicates an expected call of ProtocolVersion.
func (mr *MockBackendProtocolManagerMockRecorder) ProtocolVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtocolVersion", reflect.TypeOf((*MockBackendProtocolManager)(nil).ProtocolVersion))
}

// ReBroadcastTxs mocks base method.
func (m *MockBackendProtocolManager) ReBroadcastTxs(arg0 types.Transactions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReBroadcastTxs", arg0)
}

// ReBroadcastTxs indicates an expected call of ReBroadcastTxs.
func (mr *MockBackendProtocolManagerMockRecorder) ReBroadcastTxs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReBroadcastTxs", reflect.TypeOf((*MockBackendProtocolManager)(nil).ReBroadcastTxs), arg0)
}

// SetAcceptTxs mocks base method.
func (m *MockBackendProtocolManager) SetAcceptTxs() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAcceptTxs")
}

// SetAcceptTxs indicates an expected call of SetAcceptTxs.
func (mr *MockBackendProtocolManagerMockRecorder) SetAcceptTxs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAcceptTxs", reflect.TypeOf((*MockBackendProtocolManager)(nil).SetAcceptTxs))
}

// SetRewardbase mocks base method.
func (m *MockBackendProtocolManager) SetRewardbase(arg0 common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRewardbase", arg0)
}

// SetRewardbase indicates an expected call of SetRewardbase.
func (mr *MockBackendProtocolManagerMockRecorder) SetRewardbase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRewardbase", reflect.TypeOf((*MockBackendProtocolManager)(nil).SetRewardbase), arg0)
}

// SetRewardbaseWallet mocks base method.
func (m *MockBackendProtocolManager) SetRewardbaseWallet(arg0 accounts.Wallet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRewardbaseWallet", arg0)
}

// SetRewardbaseWallet indicates an expected call of SetRewardbaseWallet.
func (mr *MockBackendProtocolManagerMockRecorder) SetRewardbaseWallet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRewardbaseWallet", reflect.TypeOf((*MockBackendProtocolManager)(nil).SetRewardbaseWallet), arg0)
}

// SetTmpStop mocks base method.
func (m *MockBackendProtocolManager) SetSyncStop(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSyncStop", arg0)
}

// SetTmpStop indicates an expected call of SetTmpStop.
func (mr *MockBackendProtocolManagerMockRecorder) SetTmpStop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSyncStop", reflect.TypeOf((*MockBackendProtocolManager)(nil).SetSyncStop), arg0)
}

// SetWsEndPoint mocks base method.
func (m *MockBackendProtocolManager) SetWsEndPoint(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetWsEndPoint", arg0)
}

// SetWsEndPoint indicates an expected call of SetWsEndPoint.
func (mr *MockBackendProtocolManagerMockRecorder) SetWsEndPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWsEndPoint", reflect.TypeOf((*MockBackendProtocolManager)(nil).SetWsEndPoint), arg0)
}

// Start mocks base method.
func (m *MockBackendProtocolManager) Start(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockBackendProtocolManagerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBackendProtocolManager)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockBackendProtocolManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBackendProtocolManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBackendProtocolManager)(nil).Stop))
}
