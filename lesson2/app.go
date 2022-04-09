package main

import (
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

var _ types.Application = (*App)(nil)

type App struct {
}

var logger, err = log.NewDefaultLogger(log.LogFormatPlain, log.LogLevelInfo, true)

func (a App) Info(info types.RequestInfo) types.ResponseInfo {
	logger.Info("Info")
	return types.ResponseInfo{}
}

func (a App) Query(query types.RequestQuery) types.ResponseQuery {
	logger.Info("Query")
	return types.ResponseQuery{}
}

func (a App) CheckTx(tx types.RequestCheckTx) types.ResponseCheckTx {
	logger.Info("CheckTx")
	return types.ResponseCheckTx{
		Code: types.CodeTypeOK,
	}
}

func (a App) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	logger.Info("InitChain")
	return types.ResponseInitChain{}
}

func (a App) BeginBlock(block types.RequestBeginBlock) types.ResponseBeginBlock {
	logger.Info("BeginBlock")
	return types.ResponseBeginBlock{}
}

func (a App) DeliverTx(tx types.RequestDeliverTx) types.ResponseDeliverTx {
	logger.Info("DeliverTx")
	return types.ResponseDeliverTx{}
}

func (a App) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	logger.Info("EndBlock")
	return types.ResponseEndBlock{}
}

func (a App) Commit() types.ResponseCommit {
	logger.Info("Commit")
	return types.ResponseCommit{}
}

func (a App) ListSnapshots(snapshots types.RequestListSnapshots) types.ResponseListSnapshots {
	logger.Info("ListSnapshots")
	return types.ResponseListSnapshots{}
}

func (a App) OfferSnapshot(snapshot types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	logger.Info("OfferSnapshot")
	return types.ResponseOfferSnapshot{}
}

func (a App) LoadSnapshotChunk(chunk types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	logger.Info("LoadSnapshotChunk")
	return types.ResponseLoadSnapshotChunk{}
}

func (a App) ApplySnapshotChunk(chunk types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	logger.Info("ApplySnapshotChunk")
	return types.ResponseApplySnapshotChunk{}
}
