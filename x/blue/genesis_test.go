package blue_test

import (
	"testing"

	keepertest "blue/testutil/keeper"
	"blue/testutil/nullify"
	"blue/x/blue"
	"blue/x/blue/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlueKeeper(t)
	blue.InitGenesis(ctx, *k, genesisState)
	got := blue.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
