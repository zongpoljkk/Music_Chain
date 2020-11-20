package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetArtist{}

type MsgSetArtist struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
}

func NewMsgSetArtist(creator sdk.AccAddress, id string, name string) MsgSetArtist {
  return MsgSetArtist{
    ID: id,
		Creator: creator,
    Name: name,
	}
}

func (msg MsgSetArtist) Route() string {
  return RouterKey
}

func (msg MsgSetArtist) Type() string {
  return "SetArtist"
}

func (msg MsgSetArtist) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetArtist) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetArtist) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}