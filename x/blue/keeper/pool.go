package keeper

import (
	"blue/x/blue/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func (k Keeper) AppendPost() uint64 {
//   count := k.GetPostCount()
//   store.Set()
//   k.SetPostCount()
//   return count
// }

func (k Keeper) GetPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PoolCountKey))
	byteKey := []byte(types.PoolCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PoolCountKey))
	byteKey := []byte(types.PoolCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendPool(ctx sdk.Context, pool types.Pool) uint64 {
	count := k.GetPoolCount(ctx)
	pool.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PoolKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, pool.Id)
	appendedValue := k.cdc.MustMarshal(&pool)
	store.Set(byteKey, appendedValue)
	k.SetPoolCount(ctx, count+1)
	return count
}
