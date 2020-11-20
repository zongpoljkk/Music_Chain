package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
)

func handleMsgSetArtist(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetArtist) (*sdk.Result, error) {
	var artist = types.Artist{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
	}
	if !msg.Creator.Equals(k.GetArtistOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetArtist(ctx, artist)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
