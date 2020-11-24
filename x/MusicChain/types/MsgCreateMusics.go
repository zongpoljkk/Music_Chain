package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/utils"
)

var _ sdk.Msg = &MsgCreateMusics{}

type MsgCreateMusics struct {
	ID        string
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	MediaLink string         `json:"mediaLink" yaml:"mediaLink"`
	Price     int32          `json:"price" yaml:"price"`
	Name      string         `json:"name" yaml:"name"`
}

// UploadFileAndCreateMsgCreateMusics upload given `data` to GCS
// and return completed MsgCreateMusics with MediaLink as UUID().`ext`
func UploadFileAndCreateMsgCreateMusics(creator sdk.AccAddress, price int32, name string, data []byte, extWithDot string) (MsgCreateMusics, error) {
	msg := MsgCreateMusics{
		ID:      uuid.New().String(),
		Creator: creator,
		Price:   price,
		Name:    name,
	}
	fileAttr, err := utils.UploadBytesToGCS(msg.ID+extWithDot, data)
	if err == nil {
		msg.MediaLink = fileAttr.Name
	}
	return msg, err
}

func (msg MsgCreateMusics) Route() string {
	return RouterKey
}

func (msg MsgCreateMusics) Type() string {
	return "CreateMusics"
}

func (msg MsgCreateMusics) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateMusics) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMusics) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if len(msg.MediaLink) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "MediaLink can't be empty")
	}
	return nil
}
