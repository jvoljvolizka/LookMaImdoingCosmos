package keeper

import (
	"context"

	"blue/x/blue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePrefPool(goCtx context.Context, msg *types.MsgCreatePrefPool) (*types.MsgCreatePrefPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var newPool = types.PrefPool{}
	id := k.AppendPrefPool(ctx, newPool)

	return &types.MsgCreatePrefPoolResponse{Id: id}, nil
}
