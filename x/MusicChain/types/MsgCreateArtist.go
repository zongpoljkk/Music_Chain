package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateArtist{}

type MsgCreateArtist struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}

func NewMsgCreateArtist(creator sdk.AccAddress, name string) MsgCreateArtist {
	return MsgCreateArtist{
		Creator: creator,
		Name:    name,
	}
}

func (msg MsgCreateArtist) Route() string {
	return RouterKey
}

func (msg MsgCreateArtist) Type() string {
	return "CreateArtist"
}

func (msg MsgCreateArtist) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateArtist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateArtist) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
