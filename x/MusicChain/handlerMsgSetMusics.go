package MusicChain

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func handleMsgSetMusics(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMusics) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetMusicsOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	musics, err := k.GetMusics(ctx, msg.ID)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	b, err := json.Marshal(msg)
	err = json.Unmarshal(b, &musics)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	k.SetMusics(ctx, musics)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
