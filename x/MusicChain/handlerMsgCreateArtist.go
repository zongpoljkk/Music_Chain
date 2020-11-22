package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func handleMsgCreateArtist(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateArtist) (*sdk.Result, error) {
	var artist = types.Artist{
		Creator: msg.Creator,
		Name:    msg.Name,
	}
	k.CreateArtist(ctx, artist)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
