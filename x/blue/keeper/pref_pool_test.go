package keeper_test

import (
	"testing"

	keepertest "blue/testutil/keeper"
	"blue/testutil/nullify"
	"blue/x/blue/keeper"
	"blue/x/blue/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPrefPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PrefPool {
	items := make([]types.PrefPool, n)
	for i := range items {
		items[i].Id = keeper.AppendPrefPool(ctx, items[i])
	}
	return items
}

func TestPrefPoolGet(t *testing.T) {
	keeper, ctx := keepertest.BlueKeeper(t)
	items := createNPrefPool(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPrefPool(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPrefPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.BlueKeeper(t)
	items := createNPrefPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrefPool(ctx, item.Id)
		_, found := keeper.GetPrefPool(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPrefPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlueKeeper(t)
	items := createNPrefPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPrefPool(ctx)),
	)
}

func TestPrefPoolCount(t *testing.T) {
	keeper, ctx := keepertest.BlueKeeper(t)
	items := createNPrefPool(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPrefPoolCount(ctx))
}
