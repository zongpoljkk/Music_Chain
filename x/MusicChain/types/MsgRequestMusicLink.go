package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestPurchasedMusicTemporarilyLink{}

type MsgRequestPurchasedMusicTemporarilyLink struct {
	Requestor sdk.AccAddress `json:"creator" yaml:"creator"`
	MusicID   string         `json:"music_id" yaml:"music_id"`
	Timestamp time.Time      `json:"ts" yaml:"ts"`
}

func ParseMsgRequestPurchasedMusicTemporarilyLink(data []byte) MsgRequestPurchasedMusicTemporarilyLink {
	var msg MsgRequestPurchasedMusicTemporarilyLink
	json.Unmarshal(data, &msg)
	return msg
}

func (msg MsgRequestPurchasedMusicTemporarilyLink) Route() string {
	return RouterKey
}

func (msg MsgRequestPurchasedMusicTemporarilyLink) Type() string {
	return "ReqPMusicTempL"
}

func (msg MsgRequestPurchasedMusicTemporarilyLink) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Requestor)}
}

func (msg MsgRequestPurchasedMusicTemporarilyLink) GetSignBytes() []byte {
	bz, _ := json.Marshal(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgRequestPurchasedMusicTemporarilyLink) ValidateBasic() error {
	if msg.Requestor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if msg.Timestamp.After(time.Now()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "timestamp (ts) can't be in advance")
	}
	return nil
}
