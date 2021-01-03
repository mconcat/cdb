package app

// store.go defines store related methods
// copy and pasted from baseapp.go

import (
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// hack to track and modify internal multistore using storeloader
func (app *DebugApp) StoreLoader() bam.StoreLoader {
	return func(ms sdk.CommitMultiStore) error {
		app.cms = ms
		return bam.DefaultStoreLoader(ms)
	}
}
