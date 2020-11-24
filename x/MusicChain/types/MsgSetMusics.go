package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMusics{}

type MsgSetMusics struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Price   int32          `json:"price" yaml:"price"`
	Name    string         `json:"name" yaml:"name"`
}

func NewMsgSetMusics(creator sdk.AccAddress, id string, price int32, name string) MsgSetMusics {
	return MsgSetMusics{
		ID:      id,
		Creator: creator,
		Price:   price,
		Name:    name,
	}
}

func (msg MsgSetMusics) Route() string {
	return RouterKey
}

func (msg MsgSetMusics) Type() string {
	return "SetMusics"
}

func (msg MsgSetMusics) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetMusics) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetMusics) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
