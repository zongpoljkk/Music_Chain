package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateArtist creates a artist
func (k Keeper) CreateArtist(ctx sdk.Context, artist types.Artist) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ArtistPrefix + artist.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(artist)
	store.Set(key, value)
}

// GetArtist returns the artist information
func (k Keeper) GetArtist(ctx sdk.Context, key string) (types.Artist, error) {
	store := ctx.KVStore(k.storeKey)
	var artist types.Artist
	byteKey := []byte(types.ArtistPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

// SetArtist sets a artist
func (k Keeper) SetArtist(ctx sdk.Context, artist types.Artist) {
	artistKey := artist.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(artist)
	key := []byte(types.ArtistPrefix + artistKey)
	store.Set(key, bz)
}

// DeleteArtist deletes a artist
func (k Keeper) DeleteArtist(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ArtistPrefix + key))
}

//
// Functions used by querier
//

func listArtist(ctx sdk.Context, k Keeper) ([]byte, error) {
	var artistList []types.Artist
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ArtistPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var artist types.Artist
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &artist)
		artistList = append(artistList, artist)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, artistList)
	return res, nil
}

func getArtist(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	artist, err := k.GetArtist(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, artist)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetArtistOwner(ctx sdk.Context, key string) sdk.AccAddress {
	artist, err := k.GetArtist(ctx, key)
	if err != nil {
		return nil
	}
	return artist.Creator
}


// Check if the key exists in the store
func (k Keeper) ArtistExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ArtistPrefix + key))
}
