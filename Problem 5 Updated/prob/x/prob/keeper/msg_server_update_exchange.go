package keeper

import (
	"context"
	"fmt"

	"prob/x/prob/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateExchange(goCtx context.Context, msg *types.MsgUpdateExchange) (*types.MsgUpdateExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.getExchangeWithID(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Sender {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	var exchange = types.Exchange{
		Date:     val.Date,
		Amount:   val.Amount,
		Message:  msg.Message,
		Reciever: val.Reciever,
		Sender:   val.Sender,
	}
	k.updateExchange(ctx, exchange)
	return &types.MsgUpdateExchangeResponse{}, nil
}
