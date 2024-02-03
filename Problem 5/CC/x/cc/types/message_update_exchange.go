package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateExchange{}

func NewMsgUpdateExchange(creator string, date string, message string, id uint64) *MsgUpdateExchange {
	return &MsgUpdateExchange{
		Creator: creator,
		Date:    date,
		Message: message,
		Id:      id,
	}
}

func (msg *MsgUpdateExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
