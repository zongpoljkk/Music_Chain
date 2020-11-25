package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zongpoljkk/MusicChain/x/MusicChain/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreatePurchased creates a purchased
func (k Keeper) CreatePurchased(ctx sdk.Context, purchased types.Purchased) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PurchasedPrefix + purchased.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(purchased)
	store.Set(key, value)
}

// GetPurchased returns the purchased information
func (k Keeper) GetPurchased(ctx sdk.Context, key string) (types.Purchased, error) {
	store := ctx.KVStore(k.storeKey)
	var purchased types.Purchased
	byteKey := []byte(types.PurchasedPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &purchased)
	if err != nil {
		return purchased, err
	}
	return purchased, nil
}

// SetPurchased sets a purchased
func (k Keeper) SetPurchased(ctx sdk.Context, purchased types.Purchased) {
	purchasedKey := purchased.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(purchased)
	key := []byte(types.PurchasedPrefix + purchasedKey)
	store.Set(key, bz)
}

// DeletePurchased deletes a purchased
func (k Keeper) DeletePurchased(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.PurchasedPrefix + key))
}

//
// Functions used by querier
//

func listPurchased(ctx sdk.Context, k Keeper) ([]byte, error) {
	var purchasedList []types.Purchased
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PurchasedPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var purchased types.Purchased
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &purchased)
		purchasedList = append(purchasedList, purchased)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, purchasedList)
	return res, nil
}

func getPurchased(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	purchased, err := k.GetPurchased(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, purchased)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPurchasedOwner(ctx sdk.Context, key string) sdk.AccAddress {
	purchased, err := k.GetPurchased(ctx, key)
	if err != nil {
		return nil
	}
	return purchased.Creator
}


// Check if the key exists in the store
func (k Keeper) PurchasedExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PurchasedPrefix + key))
}
