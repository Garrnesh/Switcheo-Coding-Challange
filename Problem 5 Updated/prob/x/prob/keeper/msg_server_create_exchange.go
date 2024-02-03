package keeper

import (
	"context"

	"prob/x/prob/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateExchange(goCtx context.Context, msg *types.MsgCreateExchange) (*types.MsgCreateExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var exchange = types.Exchange{
		Date:     msg.Date,
		Amount:   msg.Amount,
		Message:  msg.Message,
		Reciever: msg.Reciever,
		Sender:   msg.Creator,
	}
	id := k.AppendExchange(
		ctx,
		exchange,
	)
	return &types.MsgCreateExchangeResponse{
		Id: id,
	}, nil
}
