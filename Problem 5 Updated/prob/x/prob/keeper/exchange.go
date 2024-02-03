package keeper

import (
	"encoding/binary"

	"prob/x/prob/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Create new block
func (k Keeper) AppendExchange(ctx sdk.Context, exchange types.Exchange) uint64 {
	count := k.Get_Total_Count(ctx)
	exchange.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExchangeKey))
	appendedValue := k.cdc.MustMarshal(&exchange)
	store.Set(convert_Byte(exchange.Id), appendedValue)
	k.setCount(ctx, count+1)
	return count
}

func (k Keeper) Get_Total_Count(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ExchangeCountKey)
	bs := store.Get(byteKey)
	if bs == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bs)
}

func convert_Byte(id uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, id)
	return bs
}

func (k Keeper) setCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	key_bytes := types.KeyPrefix(types.ExchangeCountKey)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, count)
	store.Set(key_bytes, bs)
}

// Update Block
func (k Keeper) getExchangeWithID(ctx sdk.Context, id uint64) (val types.Exchange, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExchangeKey))
	exchange_object := store.Get(convert_Byte(id))
	if exchange_object == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(exchange_object, &val)
	return val, true
}

func (k Keeper) updateExchange(ctx sdk.Context, exchange types.Exchange) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExchangeKey))
	ex := k.cdc.MustMarshal(&exchange)
	store.Set(convert_Byte(exchange.Id), ex)
}

// Remove Block
func (k Keeper) deleteExchange(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExchangeKey))
	store.Delete(convert_Byte(id))
}
