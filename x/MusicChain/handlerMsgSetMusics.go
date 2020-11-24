package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func handleMsgSetMusics(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMusics) (*sdk.Result, error) {
	var musics = types.Musics{
		Creator:   msg.Creator,
		ID:        msg.ID,
		MediaLink: msg.MediaLink,
		Price:     msg.Price,
		Name:      msg.Name,
	}
	if !msg.Creator.Equals(k.GetMusicsOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetMusics(ctx, musics)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
