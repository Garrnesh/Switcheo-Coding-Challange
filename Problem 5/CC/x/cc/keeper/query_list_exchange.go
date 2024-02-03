package keeper

import (
	"context"

	"CC/x/cc/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListExchange(ctx context.Context, req *types.QueryListExchangeRequest) (*types.QueryListExchangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExchangeKey))

	var exchanges []types.Exchange
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var exchange types.Exchange
		if err := k.cdc.Unmarshal(value, &exchange); err != nil {
			return err
		}

		exchanges = append(exchanges, exchange)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListExchangeResponse{Exchange: exchanges, Pagination: pageRes}, nil
}
