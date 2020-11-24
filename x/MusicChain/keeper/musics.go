package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
)

// CreateMusics creates a musics
func (k Keeper) CreateMusics(ctx sdk.Context, musics types.Musics) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MusicsPrefix + musics.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(musics)
	store.Set(key, value)
}

// GetMusics returns the musics information
func (k Keeper) GetMusics(ctx sdk.Context, key string) (types.Musics, error) {
	store := ctx.KVStore(k.storeKey)
	var musics types.Musics
	byteKey := []byte(types.MusicsPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &musics)
	if err != nil {
		return musics, err
	}
	return musics, nil
}

// SetMusics sets a musics
func (k Keeper) SetMusics(ctx sdk.Context, musics types.Musics) {
	musicsKey := musics.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(musics)
	key := []byte(types.MusicsPrefix + musicsKey)
	store.Set(key, bz)
}

// DeleteMusics deletes a musics
func (k Keeper) DeleteMusics(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MusicsPrefix + key))
}

//
// Functions used by querier
//

func listMusics(ctx sdk.Context, k Keeper) ([]byte, error) {
	var musicsList []types.Musics
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MusicsPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var musics types.Musics
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &musics)
		musicsList = append(musicsList, musics)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, musicsList)
	return res, nil
}

func getMusics(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	musics, err := k.GetMusics(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, musics)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetMusicsOwner(ctx sdk.Context, key string) sdk.AccAddress {
	musics, err := k.GetMusics(ctx, key)
	if err != nil {
		return nil
	}
	return musics.Creator
}

// Check if the key exists in the store
func (k Keeper) MusicsExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MusicsPrefix + key))
}
