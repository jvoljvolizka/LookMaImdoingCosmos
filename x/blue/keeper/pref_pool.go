package keeper

import (
	"encoding/binary"

	"blue/x/blue/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPrefPoolCount get the total number of prefPool
func (k Keeper) GetPrefPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PrefPoolCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPrefPoolCount set the total number of prefPool
func (k Keeper) SetPrefPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PrefPoolCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPrefPool appends a prefPool in the store with a new id and update the count
func (k Keeper) AppendPrefPool(
	ctx sdk.Context,
	prefPool types.PrefPool,
) uint64 {
	// Create the prefPool
	count := k.GetPrefPoolCount(ctx)

	// Set the ID of the appended value
	prefPool.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefPoolKey))
	appendedValue := k.cdc.MustMarshal(&prefPool)
	store.Set(GetPrefPoolIDBytes(prefPool.Id), appendedValue)

	// Update prefPool count
	k.SetPrefPoolCount(ctx, count+1)

	return count
}

// SetPrefPool set a specific prefPool in the store
func (k Keeper) SetPrefPool(ctx sdk.Context, prefPool types.PrefPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefPoolKey))
	b := k.cdc.MustMarshal(&prefPool)
	store.Set(GetPrefPoolIDBytes(prefPool.Id), b)
}

// GetPrefPool returns a prefPool from its id
func (k Keeper) GetPrefPool(ctx sdk.Context, id uint64) (val types.PrefPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefPoolKey))
	b := store.Get(GetPrefPoolIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePrefPool removes a prefPool from the store
func (k Keeper) RemovePrefPool(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefPoolKey))
	store.Delete(GetPrefPoolIDBytes(id))
}

// GetAllPrefPool returns all prefPool
func (k Keeper) GetAllPrefPool(ctx sdk.Context) (list []types.PrefPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefPoolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PrefPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPrefPoolIDBytes returns the byte representation of the ID
func GetPrefPoolIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPrefPoolIDFromBytes returns ID in uint64 format from a byte array
func GetPrefPoolIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
