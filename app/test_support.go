package app

import (
	ibckeeper "github.com/cosmos/ibc-go/v10/modules/core/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

func (app *WasmApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

func (app *WasmApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *WasmApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *WasmApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return app.StakingKeeper
}

func (app *WasmApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}

func (app *WasmApp) GetWasmKeeper() wasmkeeper.Keeper {
	return app.WasmKeeper
}
