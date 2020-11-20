package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
)

// Handle a message to delete name
func handleMsgDeleteArtist(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteArtist) (*sdk.Result, error) {
	if !k.ArtistExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetArtistOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteArtist(ctx, msg.ID)
	return &sdk.Result{}, nil
}
