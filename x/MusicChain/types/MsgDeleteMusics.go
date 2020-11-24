package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMusics{}

type MsgDeleteMusics struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteMusics(id string, creator sdk.AccAddress) MsgDeleteMusics {
	return MsgDeleteMusics{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgDeleteMusics) Route() string {
	return RouterKey
}

func (msg MsgDeleteMusics) Type() string {
	return "DeleteMusics"
}

func (msg MsgDeleteMusics) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteMusics) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMusics) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
