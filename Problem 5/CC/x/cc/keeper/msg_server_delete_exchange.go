package keeper

import (
	"context"
	"fmt"

	"CC/x/cc/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteExchange(goCtx context.Context, msg *types.MsgDeleteExchange) (*types.MsgDeleteExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.getExchangeWithID(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Sender {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.deleteExchange(ctx, msg.Id)
	return &types.MsgDeleteExchangeResponse{}, nil
}
