package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateExchange{}

func NewMsgCreateExchange(creator string, date string, amount string, message string, reciever string) *MsgCreateExchange {
	return &MsgCreateExchange{
		Creator:  creator,
		Date:     date,
		Amount:   amount,
		Message:  message,
		Reciever: reciever,
	}
}

func (msg *MsgCreateExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
