package keeper

import (
	"context"
	"encoding/binary"

	"blue/x/blue/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Showpool(goCtx context.Context, req *types.QueryShowpoolRequest) (*types.QueryShowpoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var pool types.Pool

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, []byte(types.PoolKey))

	byteKey := make([]byte, 8)

	binary.BigEndian.PutUint64(byteKey, req.Id)

	bz := poolStore.Get(byteKey)
	k.cdc.MustUnmarshal(bz, &pool)

	return &types.QueryShowpoolResponse{Pool: &pool}, nil
}
