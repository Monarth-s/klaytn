// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/klaytn/klaytn/work (interfaces: BlockChain)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	big "math/big"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	blockchain "github.com/klaytn/klaytn/blockchain"
	state "github.com/klaytn/klaytn/blockchain/state"
	types "github.com/klaytn/klaytn/blockchain/types"
	vm "github.com/klaytn/klaytn/blockchain/vm"
	common "github.com/klaytn/klaytn/common"
	consensus "github.com/klaytn/klaytn/consensus"
	event "github.com/klaytn/klaytn/event"
	params "github.com/klaytn/klaytn/params"
	rlp "github.com/klaytn/klaytn/rlp"
)

// MockBlockChain is a mock of BlockChain interface.
type MockBlockChain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockChainMockRecorder
}

// MockBlockChainMockRecorder is the mock recorder for MockBlockChain.
type MockBlockChainMockRecorder struct {
	mock *MockBlockChain
}

// NewMockBlockChain creates a new mock instance.
func NewMockBlockChain(ctrl *gomock.Controller) *MockBlockChain {
	mock := &MockBlockChain{ctrl: ctrl}
	mock.recorder = &MockBlockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockChain) EXPECT() *MockBlockChainMockRecorder {
	return m.recorder
}

// ApplyTransaction mocks base method.
func (m *MockBlockChain) ApplyTransaction(arg0 *params.ChainConfig, arg1 *common.Address, arg2 *state.StateDB, arg3 *types.Header, arg4 *types.Transaction, arg5 *uint64, arg6 *vm.Config) (*types.Receipt, uint64, *vm.InternalTxTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTransaction", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(*vm.InternalTxTrace)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ApplyTransaction indicates an expected call of ApplyTransaction.
func (mr *MockBlockChainMockRecorder) ApplyTransaction(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTransaction", reflect.TypeOf((*MockBlockChain)(nil).ApplyTransaction), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// BadBlocks mocks base method.
func (m *MockBlockChain) BadBlocks() ([]blockchain.BadBlockArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BadBlocks")
	ret0, _ := ret[0].([]blockchain.BadBlockArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BadBlocks indicates an expected call of BadBlocks.
func (mr *MockBlockChainMockRecorder) BadBlocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BadBlocks", reflect.TypeOf((*MockBlockChain)(nil).BadBlocks))
}

// BlockSubscriptionLoop mocks base method.
func (m *MockBlockChain) BlockSubscriptionLoop(arg0 *blockchain.TxPool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockSubscriptionLoop", arg0)
}

// BlockSubscriptionLoop indicates an expected call of BlockSubscriptionLoop.
func (mr *MockBlockChainMockRecorder) BlockSubscriptionLoop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSubscriptionLoop", reflect.TypeOf((*MockBlockChain)(nil).BlockSubscriptionLoop), arg0)
}

// CloseBlockSubscriptionLoop mocks base method.
func (m *MockBlockChain) CloseBlockSubscriptionLoop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseBlockSubscriptionLoop")
}

// CloseBlockSubscriptionLoop indicates an expected call of CloseBlockSubscriptionLoop.
func (mr *MockBlockChainMockRecorder) CloseBlockSubscriptionLoop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseBlockSubscriptionLoop", reflect.TypeOf((*MockBlockChain)(nil).CloseBlockSubscriptionLoop))
}

// Config mocks base method.
func (m *MockBlockChain) Config() *params.ChainConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*params.ChainConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockBlockChainMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBlockChain)(nil).Config))
}

// ContractCodeWithPrefix indicates an expected call of ContractCodeWithPrefix
func (mr *MockBlockChainMockRecorder) ContractCodeWithPrefix(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContractCodeWithPrefix", reflect.TypeOf((*MockBlockChain)(nil).ContractCodeWithPrefix), arg0)
}

// CurrentBlock mocks base method
func (m *MockBlockChain) CurrentBlock() *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// CurrentBlock indicates an expected call of CurrentBlock.
func (mr *MockBlockChainMockRecorder) CurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockBlockChain)(nil).CurrentBlock))
}

// CurrentFastBlock mocks base method.
func (m *MockBlockChain) CurrentFastBlock() *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentFastBlock")
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// CurrentFastBlock indicates an expected call of CurrentFastBlock.
func (mr *MockBlockChainMockRecorder) CurrentFastBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentFastBlock", reflect.TypeOf((*MockBlockChain)(nil).CurrentFastBlock))
}

// CurrentHeader mocks base method.
func (m *MockBlockChain) CurrentHeader() *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentHeader")
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// CurrentHeader indicates an expected call of CurrentHeader.
func (mr *MockBlockChainMockRecorder) CurrentHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockBlockChain)(nil).CurrentHeader))
}

// Engine mocks base method.
func (m *MockBlockChain) Engine() consensus.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Engine")
	ret0, _ := ret[0].(consensus.Engine)
	return ret0
}

// Engine indicates an expected call of Engine.
func (mr *MockBlockChainMockRecorder) Engine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Engine", reflect.TypeOf((*MockBlockChain)(nil).Engine))
}

// Export mocks base method.
func (m *MockBlockChain) Export(arg0 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Export indicates an expected call of Export.
func (mr *MockBlockChainMockRecorder) Export(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockBlockChain)(nil).Export), arg0)
}

// FastSyncCommitHead mocks base method.
func (m *MockBlockChain) FastSyncCommitHead(arg0 common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FastSyncCommitHead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FastSyncCommitHead indicates an expected call of FastSyncCommitHead.
func (mr *MockBlockChainMockRecorder) FastSyncCommitHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FastSyncCommitHead", reflect.TypeOf((*MockBlockChain)(nil).FastSyncCommitHead), arg0)
}

// Genesis mocks base method.
func (m *MockBlockChain) Genesis() *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// Genesis indicates an expected call of Genesis.
func (mr *MockBlockChainMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockBlockChain)(nil).Genesis))
}

// GetBlock mocks base method.
func (m *MockBlockChain) GetBlock(arg0 common.Hash, arg1 uint64) *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockBlockChainMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockBlockChain)(nil).GetBlock), arg0, arg1)
}

// GetBlockByHash mocks base method.
func (m *MockBlockChain) GetBlockByHash(arg0 common.Hash) *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// GetBlockByHash indicates an expected call of GetBlockByHash.
func (mr *MockBlockChainMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockBlockChain)(nil).GetBlockByHash), arg0)
}

// GetBlockByNumber mocks base method.
func (m *MockBlockChain) GetBlockByNumber(arg0 uint64) *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber.
func (mr *MockBlockChainMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockBlockChain)(nil).GetBlockByNumber), arg0)
}

// GetBlockHashesFromHash mocks base method.
func (m *MockBlockChain) GetBlockHashesFromHash(arg0 common.Hash, arg1 uint64) []common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHashesFromHash", arg0, arg1)
	ret0, _ := ret[0].([]common.Hash)
	return ret0
}

// GetBlockHashesFromHash indicates an expected call of GetBlockHashesFromHash.
func (mr *MockBlockChainMockRecorder) GetBlockHashesFromHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHashesFromHash", reflect.TypeOf((*MockBlockChain)(nil).GetBlockHashesFromHash), arg0, arg1)
}

// GetBlockReceiptsInCache mocks base method.
func (m *MockBlockChain) GetBlockReceiptsInCache(arg0 common.Hash) types.Receipts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceiptsInCache", arg0)
	ret0, _ := ret[0].(types.Receipts)
	return ret0
}

// GetBlockReceiptsInCache indicates an expected call of GetBlockReceiptsInCache.
func (mr *MockBlockChainMockRecorder) GetBlockReceiptsInCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceiptsInCache", reflect.TypeOf((*MockBlockChain)(nil).GetBlockReceiptsInCache), arg0)
}

// GetBodyRLP mocks base method.
func (m *MockBlockChain) GetBodyRLP(arg0 common.Hash) rlp.RawValue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBodyRLP", arg0)
	ret0, _ := ret[0].(rlp.RawValue)
	return ret0
}

// GetBodyRLP indicates an expected call of GetBodyRLP.
func (mr *MockBlockChainMockRecorder) GetBodyRLP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodyRLP", reflect.TypeOf((*MockBlockChain)(nil).GetBodyRLP), arg0)
}

// GetContractStorageRoot mocks base method.
func (m *MockBlockChain) GetContractStorageRoot(arg0 *types.Block, arg1 state.Database, arg2 common.Address) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractStorageRoot", arg0, arg1, arg2)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractStorageRoot indicates an expected call of GetContractStorageRoot.
func (mr *MockBlockChainMockRecorder) GetContractStorageRoot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractStorageRoot", reflect.TypeOf((*MockBlockChain)(nil).GetContractStorageRoot), arg0, arg1, arg2)
}

// GetHeader mocks base method.
func (m *MockBlockChain) GetHeader(arg0 common.Hash, arg1 uint64) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockBlockChainMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockBlockChain)(nil).GetHeader), arg0, arg1)
}

// GetHeaderByHash mocks base method.
func (m *MockBlockChain) GetHeaderByHash(arg0 common.Hash) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", arg0)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash.
func (mr *MockBlockChainMockRecorder) GetHeaderByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockBlockChain)(nil).GetHeaderByHash), arg0)
}

// GetHeaderByNumber mocks base method.
func (m *MockBlockChain) GetHeaderByNumber(arg0 uint64) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByNumber", arg0)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeaderByNumber indicates an expected call of GetHeaderByNumber.
func (mr *MockBlockChainMockRecorder) GetHeaderByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByNumber", reflect.TypeOf((*MockBlockChain)(nil).GetHeaderByNumber), arg0)
}

// GetLogsByHash mocks base method.
func (m *MockBlockChain) GetLogsByHash(arg0 common.Hash) [][]*types.Log {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogsByHash", arg0)
	ret0, _ := ret[0].([][]*types.Log)
	return ret0
}

// GetLogsByHash indicates an expected call of GetLogsByHash.
func (mr *MockBlockChainMockRecorder) GetLogsByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsByHash", reflect.TypeOf((*MockBlockChain)(nil).GetLogsByHash), arg0)
}

// GetReceiptsByBlockHash mocks base method.
func (m *MockBlockChain) GetReceiptsByBlockHash(arg0 common.Hash) types.Receipts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptsByBlockHash", arg0)
	ret0, _ := ret[0].(types.Receipts)
	return ret0
}

// GetReceiptsByBlockHash indicates an expected call of GetReceiptsByBlockHash.
func (mr *MockBlockChainMockRecorder) GetReceiptsByBlockHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptsByBlockHash", reflect.TypeOf((*MockBlockChain)(nil).GetReceiptsByBlockHash), arg0)
}

// GetTd mocks base method.
func (m *MockBlockChain) GetTd(arg0 common.Hash, arg1 uint64) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTd", arg0, arg1)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetTd indicates an expected call of GetTd.
func (mr *MockBlockChainMockRecorder) GetTd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTd", reflect.TypeOf((*MockBlockChain)(nil).GetTd), arg0, arg1)
}

// GetTdByHash mocks base method.
func (m *MockBlockChain) GetTdByHash(arg0 common.Hash) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTdByHash", arg0)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetTdByHash indicates an expected call of GetTdByHash.
func (mr *MockBlockChainMockRecorder) GetTdByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTdByHash", reflect.TypeOf((*MockBlockChain)(nil).GetTdByHash), arg0)
}

// GetTxAndLookupInfo mocks base method.
func (m *MockBlockChain) GetTxAndLookupInfo(arg0 common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxAndLookupInfo", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(common.Hash)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	return ret0, ret1, ret2, ret3
}

// GetTxAndLookupInfo indicates an expected call of GetTxAndLookupInfo.
func (mr *MockBlockChainMockRecorder) GetTxAndLookupInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxAndLookupInfo", reflect.TypeOf((*MockBlockChain)(nil).GetTxAndLookupInfo), arg0)
}

// GetTxAndLookupInfoInCache mocks base method.
func (m *MockBlockChain) GetTxAndLookupInfoInCache(arg0 common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxAndLookupInfoInCache", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(common.Hash)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	return ret0, ret1, ret2, ret3
}

// GetTxAndLookupInfoInCache indicates an expected call of GetTxAndLookupInfoInCache.
func (mr *MockBlockChainMockRecorder) GetTxAndLookupInfoInCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxAndLookupInfoInCache", reflect.TypeOf((*MockBlockChain)(nil).GetTxAndLookupInfoInCache), arg0)
}

// GetTxLookupInfoAndReceipt mocks base method.
func (m *MockBlockChain) GetTxLookupInfoAndReceipt(arg0 common.Hash) (*types.Transaction, common.Hash, uint64, uint64, *types.Receipt) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxLookupInfoAndReceipt", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(common.Hash)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	ret4, _ := ret[4].(*types.Receipt)
	return ret0, ret1, ret2, ret3, ret4
}

// GetTxLookupInfoAndReceipt indicates an expected call of GetTxLookupInfoAndReceipt.
func (mr *MockBlockChainMockRecorder) GetTxLookupInfoAndReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxLookupInfoAndReceipt", reflect.TypeOf((*MockBlockChain)(nil).GetTxLookupInfoAndReceipt), arg0)
}

// GetTxLookupInfoAndReceiptInCache mocks base method.
func (m *MockBlockChain) GetTxLookupInfoAndReceiptInCache(arg0 common.Hash) (*types.Transaction, common.Hash, uint64, uint64, *types.Receipt) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxLookupInfoAndReceiptInCache", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(common.Hash)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	ret4, _ := ret[4].(*types.Receipt)
	return ret0, ret1, ret2, ret3, ret4
}

// GetTxLookupInfoAndReceiptInCache indicates an expected call of GetTxLookupInfoAndReceiptInCache.
func (mr *MockBlockChainMockRecorder) GetTxLookupInfoAndReceiptInCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxLookupInfoAndReceiptInCache", reflect.TypeOf((*MockBlockChain)(nil).GetTxLookupInfoAndReceiptInCache), arg0)
}

// HasBadBlock mocks base method.
func (m *MockBlockChain) HasBadBlock(arg0 common.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBadBlock", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBadBlock indicates an expected call of HasBadBlock.
func (mr *MockBlockChainMockRecorder) HasBadBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBadBlock", reflect.TypeOf((*MockBlockChain)(nil).HasBadBlock), arg0)
}

// HasBlock mocks base method.
func (m *MockBlockChain) HasBlock(arg0 common.Hash, arg1 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBlock", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBlock indicates an expected call of HasBlock.
func (mr *MockBlockChainMockRecorder) HasBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBlock", reflect.TypeOf((*MockBlockChain)(nil).HasBlock), arg0, arg1)
}

// HasHeader mocks base method.
func (m *MockBlockChain) HasHeader(arg0 common.Hash, arg1 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasHeader", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasHeader indicates an expected call of HasHeader.
func (mr *MockBlockChainMockRecorder) HasHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasHeader", reflect.TypeOf((*MockBlockChain)(nil).HasHeader), arg0, arg1)
}

// InsertChain mocks base method.
func (m *MockBlockChain) InsertChain(arg0 types.Blocks) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertChain", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertChain indicates an expected call of InsertChain.
func (mr *MockBlockChainMockRecorder) InsertChain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertChain", reflect.TypeOf((*MockBlockChain)(nil).InsertChain), arg0)
}

// InsertHeaderChain mocks base method.
func (m *MockBlockChain) InsertHeaderChain(arg0 []*types.Header, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeaderChain", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeaderChain indicates an expected call of InsertHeaderChain.
func (mr *MockBlockChainMockRecorder) InsertHeaderChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeaderChain", reflect.TypeOf((*MockBlockChain)(nil).InsertHeaderChain), arg0, arg1)
}

// InsertReceiptChain mocks base method.
func (m *MockBlockChain) InsertReceiptChain(arg0 types.Blocks, arg1 []types.Receipts) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertReceiptChain", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertReceiptChain indicates an expected call of InsertReceiptChain.
func (mr *MockBlockChainMockRecorder) InsertReceiptChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertReceiptChain", reflect.TypeOf((*MockBlockChain)(nil).InsertReceiptChain), arg0, arg1)
}

// IsParallelDBWrite mocks base method.
func (m *MockBlockChain) IsParallelDBWrite() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsParallelDBWrite")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsParallelDBWrite indicates an expected call of IsParallelDBWrite.
func (mr *MockBlockChainMockRecorder) IsParallelDBWrite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsParallelDBWrite", reflect.TypeOf((*MockBlockChain)(nil).IsParallelDBWrite))
}

// IsSenderTxHashIndexingEnabled mocks base method.
func (m *MockBlockChain) IsSenderTxHashIndexingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSenderTxHashIndexingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSenderTxHashIndexingEnabled indicates an expected call of IsSenderTxHashIndexingEnabled.
func (mr *MockBlockChainMockRecorder) IsSenderTxHashIndexingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSenderTxHashIndexingEnabled", reflect.TypeOf((*MockBlockChain)(nil).IsSenderTxHashIndexingEnabled))
}

// PostChainEvents mocks base method.
func (m *MockBlockChain) PostChainEvents(arg0 []interface{}, arg1 []*types.Log) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostChainEvents", arg0, arg1)
}

// PostChainEvents indicates an expected call of PostChainEvents.
func (mr *MockBlockChainMockRecorder) PostChainEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostChainEvents", reflect.TypeOf((*MockBlockChain)(nil).PostChainEvents), arg0, arg1)
}

// PrepareStateMigration mocks base method.
func (m *MockBlockChain) PrepareStateMigration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareStateMigration")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareStateMigration indicates an expected call of PrepareStateMigration.
func (mr *MockBlockChainMockRecorder) PrepareStateMigration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStateMigration", reflect.TypeOf((*MockBlockChain)(nil).PrepareStateMigration))
}

// Processor mocks base method.
func (m *MockBlockChain) Processor() blockchain.Processor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Processor")
	ret0, _ := ret[0].(blockchain.Processor)
	return ret0
}

// Processor indicates an expected call of Processor.
func (mr *MockBlockChainMockRecorder) Processor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Processor", reflect.TypeOf((*MockBlockChain)(nil).Processor))
}

// ResetWithGenesisBlock mocks base method.
func (m *MockBlockChain) ResetWithGenesisBlock(arg0 *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWithGenesisBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetWithGenesisBlock indicates an expected call of ResetWithGenesisBlock.
func (mr *MockBlockChainMockRecorder) ResetWithGenesisBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWithGenesisBlock", reflect.TypeOf((*MockBlockChain)(nil).ResetWithGenesisBlock), arg0)
}

// Rollback mocks base method.
func (m *MockBlockChain) Rollback(arg0 []common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", arg0)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockBlockChainMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockBlockChain)(nil).Rollback), arg0)
}

// SaveTrieNodeCacheToDisk mocks base method.
func (m *MockBlockChain) SaveTrieNodeCacheToDisk() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTrieNodeCacheToDisk")
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTrieNodeCacheToDisk indicates an expected call of SaveTrieNodeCacheToDisk.
func (mr *MockBlockChainMockRecorder) SaveTrieNodeCacheToDisk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTrieNodeCacheToDisk", reflect.TypeOf((*MockBlockChain)(nil).SaveTrieNodeCacheToDisk))
}

// SetHead mocks base method.
func (m *MockBlockChain) SetHead(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHead indicates an expected call of SetHead.
func (mr *MockBlockChainMockRecorder) SetHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockBlockChain)(nil).SetHead), arg0)
}

// SetProposerPolicy mocks base method.
func (m *MockBlockChain) SetProposerPolicy(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProposerPolicy", arg0)
}

// SetProposerPolicy indicates an expected call of SetProposerPolicy.
func (mr *MockBlockChainMockRecorder) SetProposerPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProposerPolicy", reflect.TypeOf((*MockBlockChain)(nil).SetProposerPolicy), arg0)
}

// SetUseGiniCoeff mocks base method.
func (m *MockBlockChain) SetUseGiniCoeff(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUseGiniCoeff", arg0)
}

// SetUseGiniCoeff indicates an expected call of SetUseGiniCoeff.
func (mr *MockBlockChainMockRecorder) SetUseGiniCoeff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUseGiniCoeff", reflect.TypeOf((*MockBlockChain)(nil).SetUseGiniCoeff), arg0)
}

// StartCollectingTrieStats mocks base method.
func (m *MockBlockChain) StartCollectingTrieStats(arg0 common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCollectingTrieStats", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartCollectingTrieStats indicates an expected call of StartCollectingTrieStats.
func (mr *MockBlockChainMockRecorder) StartCollectingTrieStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCollectingTrieStats", reflect.TypeOf((*MockBlockChain)(nil).StartCollectingTrieStats), arg0)
}

// StartContractWarmUp mocks base method.
func (m *MockBlockChain) StartContractWarmUp(arg0 common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContractWarmUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContractWarmUp indicates an expected call of StartContractWarmUp.
func (mr *MockBlockChainMockRecorder) StartContractWarmUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContractWarmUp", reflect.TypeOf((*MockBlockChain)(nil).StartContractWarmUp), arg0)
}

// StartStateMigration mocks base method.
func (m *MockBlockChain) StartStateMigration(arg0 uint64, arg1 common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartStateMigration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartStateMigration indicates an expected call of StartStateMigration.
func (mr *MockBlockChainMockRecorder) StartStateMigration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartStateMigration", reflect.TypeOf((*MockBlockChain)(nil).StartStateMigration), arg0, arg1)
}

// StartWarmUp mocks base method.
func (m *MockBlockChain) StartWarmUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWarmUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWarmUp indicates an expected call of StartWarmUp.
func (mr *MockBlockChainMockRecorder) StartWarmUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWarmUp", reflect.TypeOf((*MockBlockChain)(nil).StartWarmUp))
}

// State mocks base method.
func (m *MockBlockChain) State() (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockBlockChainMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockBlockChain)(nil).State))
}

// StateAt mocks base method.
func (m *MockBlockChain) StateAt(arg0 common.Hash) (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAt", arg0)
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAt indicates an expected call of StateAt.
func (mr *MockBlockChainMockRecorder) StateAt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAt", reflect.TypeOf((*MockBlockChain)(nil).StateAt), arg0)
}

// StateAtWithGCLock mocks base method.
func (m *MockBlockChain) StateAtWithGCLock(arg0 common.Hash) (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtWithGCLock", arg0)
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAtWithGCLock indicates an expected call of StateAtWithGCLock.
func (mr *MockBlockChainMockRecorder) StateAtWithGCLock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtWithGCLock", reflect.TypeOf((*MockBlockChain)(nil).StateAtWithGCLock), arg0)
}

// StateAtWithPersistent mocks base method.
func (m *MockBlockChain) StateAtWithPersistent(arg0 common.Hash) (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtWithPersistent", arg0)
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAtWithPersistent indicates an expected call of StateAtWithPersistent.
func (mr *MockBlockChainMockRecorder) StateAtWithPersistent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtWithPersistent", reflect.TypeOf((*MockBlockChain)(nil).StateAtWithPersistent), arg0)
}

// StateCache mocks base method.
func (m *MockBlockChain) StateCache() state.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateCache")
	ret0, _ := ret[0].(state.Database)
	return ret0
}

// StateCache indicates an expected call of StateCache.
func (mr *MockBlockChainMockRecorder) StateCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateCache", reflect.TypeOf((*MockBlockChain)(nil).StateCache))
}

// StateMigrationStatus mocks base method.
func (m *MockBlockChain) StateMigrationStatus() (bool, uint64, int, int, int, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMigrationStatus")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(int)
	ret4, _ := ret[4].(int)
	ret5, _ := ret[5].(float64)
	ret6, _ := ret[6].(error)
	return ret0, ret1, ret2, ret3, ret4, ret5, ret6
}

// StateMigrationStatus indicates an expected call of StateMigrationStatus.
func (mr *MockBlockChainMockRecorder) StateMigrationStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMigrationStatus", reflect.TypeOf((*MockBlockChain)(nil).StateMigrationStatus))
}

// Stop mocks base method.
func (m *MockBlockChain) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBlockChainMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockChain)(nil).Stop))
}

// StopStateMigration mocks base method.
func (m *MockBlockChain) StopStateMigration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopStateMigration")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopStateMigration indicates an expected call of StopStateMigration.
func (mr *MockBlockChainMockRecorder) StopStateMigration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopStateMigration", reflect.TypeOf((*MockBlockChain)(nil).StopStateMigration))
}

// StopWarmUp mocks base method.
func (m *MockBlockChain) StopWarmUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopWarmUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopWarmUp indicates an expected call of StopWarmUp.
func (mr *MockBlockChainMockRecorder) StopWarmUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopWarmUp", reflect.TypeOf((*MockBlockChain)(nil).StopWarmUp))
}

// SubscribeChainEvent mocks base method.
func (m *MockBlockChain) SubscribeChainEvent(arg0 chan<- blockchain.ChainEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChainEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeChainEvent indicates an expected call of SubscribeChainEvent.
func (mr *MockBlockChainMockRecorder) SubscribeChainEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChainEvent", reflect.TypeOf((*MockBlockChain)(nil).SubscribeChainEvent), arg0)
}

// SubscribeChainHeadEvent mocks base method.
func (m *MockBlockChain) SubscribeChainHeadEvent(arg0 chan<- blockchain.ChainHeadEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChainHeadEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeChainHeadEvent indicates an expected call of SubscribeChainHeadEvent.
func (mr *MockBlockChainMockRecorder) SubscribeChainHeadEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChainHeadEvent", reflect.TypeOf((*MockBlockChain)(nil).SubscribeChainHeadEvent), arg0)
}

// SubscribeChainSideEvent mocks base method.
func (m *MockBlockChain) SubscribeChainSideEvent(arg0 chan<- blockchain.ChainSideEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChainSideEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeChainSideEvent indicates an expected call of SubscribeChainSideEvent.
func (mr *MockBlockChainMockRecorder) SubscribeChainSideEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChainSideEvent", reflect.TypeOf((*MockBlockChain)(nil).SubscribeChainSideEvent), arg0)
}

// SubscribeLogsEvent mocks base method.
func (m *MockBlockChain) SubscribeLogsEvent(arg0 chan<- []*types.Log) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeLogsEvent indicates an expected call of SubscribeLogsEvent.
func (mr *MockBlockChainMockRecorder) SubscribeLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeLogsEvent", reflect.TypeOf((*MockBlockChain)(nil).SubscribeLogsEvent), arg0)
}

// SubscribeRemovedLogsEvent mocks base method.
func (m *MockBlockChain) SubscribeRemovedLogsEvent(arg0 chan<- blockchain.RemovedLogsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeRemovedLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeRemovedLogsEvent indicates an expected call of SubscribeRemovedLogsEvent.
func (mr *MockBlockChainMockRecorder) SubscribeRemovedLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeRemovedLogsEvent", reflect.TypeOf((*MockBlockChain)(nil).SubscribeRemovedLogsEvent), arg0)
}

// TrieNode mocks base method.
func (m *MockBlockChain) TrieNode(arg0 common.Hash) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrieNode", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrieNode indicates an expected call of TrieNode.
func (mr *MockBlockChainMockRecorder) TrieNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrieNode", reflect.TypeOf((*MockBlockChain)(nil).TrieNode), arg0)
}

// Validator mocks base method.
func (m *MockBlockChain) Validator() blockchain.Validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator")
	ret0, _ := ret[0].(blockchain.Validator)
	return ret0
}

// Validator indicates an expected call of Validator.
func (mr *MockBlockChainMockRecorder) Validator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockBlockChain)(nil).Validator))
}

// WriteBlockWithState mocks base method.
func (m *MockBlockChain) WriteBlockWithState(arg0 *types.Block, arg1 []*types.Receipt, arg2 *state.StateDB) (blockchain.WriteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteBlockWithState", arg0, arg1, arg2)
	ret0, _ := ret[0].(blockchain.WriteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteBlockWithState indicates an expected call of WriteBlockWithState.
func (mr *MockBlockChainMockRecorder) WriteBlockWithState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBlockWithState", reflect.TypeOf((*MockBlockChain)(nil).WriteBlockWithState), arg0, arg1, arg2)
}
