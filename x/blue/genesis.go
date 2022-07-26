package blue

import (
	"blue/x/blue/keeper"
	"blue/x/blue/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the prefPool
	for _, elem := range genState.PrefPoolList {
		k.SetPrefPool(ctx, elem)
	}

	// Set prefPool count
	k.SetPrefPoolCount(ctx, genState.PrefPoolCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PrefPoolList = k.GetAllPrefPool(ctx)
	genesis.PrefPoolCount = k.GetPrefPoolCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
