package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func handleMsgCreateMusics(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateMusics) (*sdk.Result, error) {
	var musics = types.Musics{
		Creator:   msg.Creator,
		ID:        msg.ID,
		MediaLink: msg.MediaLink,
		Price:     msg.Price,
		Name:      msg.Name,
	}
	k.CreateMusics(ctx, musics)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
