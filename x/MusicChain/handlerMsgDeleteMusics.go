package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

// Handle a message to delete name
func handleMsgDeleteMusics(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteMusics) (*sdk.Result, error) {
	if !k.MusicsExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetMusicsOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteMusics(ctx, msg.ID)
	return &sdk.Result{}, nil
}
