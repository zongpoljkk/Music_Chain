package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePurchased{}

type MsgCreatePurchased struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	MusicID string         `json:"MusicID" yaml:"MusicID"`
}

func NewMsgCreatePurchased(creator sdk.AccAddress, MusicID string) MsgCreatePurchased {
	return MsgCreatePurchased{
		ID:      creator.String() + "-" + MusicID,
		Creator: creator,
		MusicID: MusicID,
	}
}

func (msg MsgCreatePurchased) Route() string {
	return RouterKey
}

func (msg MsgCreatePurchased) Type() string {
	return "CreatePurchased"
}

func (msg MsgCreatePurchased) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePurchased) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePurchased) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
