package app

// app.go
// wraps any application that exposes ABCI methods exposed by baseapp

import (
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	store "github.com/cosmos/cosmos-sdk/store"
)


type ABCIStage uint

const (
	StageInitChain ABCIStage = iota
	StageBeginBlock
	StageTx
	StageEndBlock
	StageCommit
)

// DebugApp replace BaseApp in applications
type DebugApp struct {
	*bam.BaseApp
	
	stage ABCIStage
	db dbm.DB
	logger log.Logger
	cms store.CommitMultiStore
	deliverState *state
}

func NewDebugApp(
	name string, logger log.Logger, db dbm.DB, txDecoder sdk.TxDecoder, options ...func(*bam.BaseApp),
) *DebugApp {
	app := &DebugApp{
		BaseApp: bam.NewBaseApp(name, logger, db, txDecoder, options...),

		stage: StageInitChain,
		db: db,
		logger: logger,
		cms: store.NewCommitMultiStore(db),
	}

	return app
}

func (app *DebugApp) setDeliverState(header abci.Header) {
	// TODO: Consider two cases: querying and writing on the latest state, querying on the previous states
	ms := app.cms.CacheMultiStore()
	app.deliverState = &state {
		ms: ms,
		ctx: sdk.NewContext(ms, header, false, app.logger),
	}
}

// Misc.

func (app *DebugApp) AssertStage(stage ABCIStage) bool {
	return app.stage == stage
}

func (app *DebugApp) NextStage() {
	if app.stage == StageCommit {
		app.stage = StageBeginBlock
	} else {
		app.stage += 1
	}
}

