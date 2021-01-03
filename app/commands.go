package app

// commands.go defines DebugApp methods corresponding to user command

func (app *DebugApp) BlockHeight() uint64 {
	
}

func (app *DebugApp) InspectState(key []byte) []byte {

}

// TODO: ModifyState(key []byte, value []byte)
// make a copy rootmultistore, modify the state using the 
// ModifyState can be called only when app.AssertStage(StageBeginBlock)
