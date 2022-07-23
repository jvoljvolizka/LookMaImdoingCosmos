package keeper_test

import (
	"context"
	"testing"

	keepertest "blue/testutil/keeper"
	"blue/x/blue/keeper"
	"blue/x/blue/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlueKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
