package keeper

import (
	"context"

	"prob/x/prob/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowExchange(goCtx context.Context, req *types.QueryShowExchangeRequest) (*types.QueryShowExchangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.getExchangeWithID(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	return &types.QueryShowExchangeResponse{Exchange: val}, nil //Changed this from &val
}
