package MusicChain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreatePurchased:
			return handleMsgCreatePurchased(ctx, k, msg)
		case types.MsgCreateMusics:
			return handleMsgCreateMusics(ctx, k, msg)
		case types.MsgSetMusics:
			return handleMsgSetMusics(ctx, k, msg)
		case types.MsgDeleteMusics:
			return handleMsgDeleteMusics(ctx, k, msg)
		case types.MsgCreateArtist:
			return handleMsgCreateArtist(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
