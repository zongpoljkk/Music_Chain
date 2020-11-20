package MusicChain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateArtist:
			return handleMsgCreateArtist(ctx, k, msg)
		case types.MsgSetArtist:
			return handleMsgSetArtist(ctx, k, msg)
		case types.MsgDeleteArtist:
			return handleMsgDeleteArtist(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
