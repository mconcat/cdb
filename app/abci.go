package app

// abci.go defines DebugApp custom ABCI methods

import (	
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ABCI commands

func (app *DebugApp) InitChain(req abci.RequestInitChain) (res abci.ResponseInitChain) {
	app.AssertStage(StageInitChain)
	initHeader := abci.Header{ChainID: req.ChainId, Time: req.Time}
	app.setDeliverState(initHeader)
	app.deliverState.ctx = app.deliverState.ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())
	return app.BaseApp.InitChain(req)
}

func (app *DebugApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	app.AssertStage(StageBeginBlock)
	
	if app.deliverState == nil {
		app.setDeliverState(req.Header)
	} else {
		app.deliverState.ctx = app.deliverState.ctx.
			WithBlockHeader(req.Header).
			WithBlockHeight(req.Header.Height)
	}

	// TODO: gas

	return app.BaseApp.BeginBlock(req)
}


