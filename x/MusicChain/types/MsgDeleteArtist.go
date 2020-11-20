package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteArtist{}

type MsgDeleteArtist struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteArtist(id string, creator sdk.AccAddress) MsgDeleteArtist {
  return MsgDeleteArtist{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteArtist) Route() string {
  return RouterKey
}

func (msg MsgDeleteArtist) Type() string {
  return "DeleteArtist"
}

func (msg MsgDeleteArtist) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteArtist) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteArtist) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}