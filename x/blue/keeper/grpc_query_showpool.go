package keeper

import (
	"context"

	"blue/x/blue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Showpool(goCtx context.Context, req *types.QueryShowpoolRequest) (*types.QueryShowpoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}


	ctx := sdk.UnwrapSDKContext(goCtx)
	pool := k.GetPool(ctx, req.Id)

	return &types.QueryShowpoolResponse{Pool: &pool}, nil
}
