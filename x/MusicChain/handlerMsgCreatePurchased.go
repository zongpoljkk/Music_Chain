package MusicChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/keeper"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

func handleMsgCreatePurchased(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePurchased) (*sdk.Result, error) {
	var purchased = types.Purchased{
		Creator: msg.Creator,
		ID:      msg.ID,
		MusicID: msg.MusicID,
	}

	if k.PurchasedExists(ctx, purchased.ID) {
		return nil, nil
	}

	music, anyFuckinError := k.GetMusics(ctx, purchased.MusicID)

	if anyFuckinError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, anyFuckinError.Error())
	}

	price := sdk.NewCoin("token", sdk.NewInt(int64(music.Price)))

	anyFuckinError = k.CoinKeeper.SendCoins(ctx, purchased.Creator, music.Creator, sdk.Coins{price})
	if anyFuckinError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, anyFuckinError.Error())
	}
	k.CreatePurchased(ctx, purchased)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
