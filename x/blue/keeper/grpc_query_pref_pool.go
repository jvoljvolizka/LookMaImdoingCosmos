package keeper

import (
	"context"

	"blue/x/blue/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PrefPoolAll(c context.Context, req *types.QueryAllPrefPoolRequest) (*types.QueryAllPrefPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var prefPools []types.PrefPool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	prefPoolStore := prefix.NewStore(store, types.KeyPrefix(types.PrefPoolKey))

	pageRes, err := query.Paginate(prefPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var prefPool types.PrefPool
		if err := k.cdc.Unmarshal(value, &prefPool); err != nil {
			return err
		}

		prefPools = append(prefPools, prefPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPrefPoolResponse{PrefPool: prefPools, Pagination: pageRes}, nil
}

func (k Keeper) PrefPool(c context.Context, req *types.QueryGetPrefPoolRequest) (*types.QueryGetPrefPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	prefPool, found := k.GetPrefPool(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPrefPoolResponse{PrefPool: prefPool}, nil
}
