package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for MusicChain clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListMusics:
			return listMusics(ctx, k)
		case types.QueryGetMusics:
			return getMusics(ctx, path[1:], k)
		case types.QueryListArtist:
			return listArtist(ctx, k)
		case types.QueryGetArtist:
			return getArtist(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown MusicChain query endpoint")
		}
	}
}
