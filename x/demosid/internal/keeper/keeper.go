package keeper

import (
	"bytes"

	"aquarelle.io/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper    bank.Keeper
	AccountKeeper auth.AccountKeeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the DemosID Keeper
func NewKeeper(coinKeeper bank.Keeper, accountKeeper auth.AccountKeeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper:    coinKeeper,
		AccountKeeper: accountKeeper,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// Gets the entire attribute metadata struct for a name
func (k Keeper) GetAttribute(ctx sdk.Context, name string, owner sdk.AccAddress) types.DemosAttribute {

	store := ctx.KVStore(k.storeKey)
	nameKey := k.AttributeStoreKey(owner, name)

	storedBytes := store.Get(nameKey)
	if storedBytes == nil {
		panic("Attribute doesn´t exists!")
	}

	var attr types.DemosAttribute
	k.cdc.MustUnmarshalBinaryBare(storedBytes, &attr)

	return attr
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteAttribute(ctx sdk.Context, attrName string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	nameKey := k.AttributeStoreKey(owner, attrName)
	store.Delete(nameKey)
}

// Changes the value for an attribute
func (k Keeper) SetValue(ctx sdk.Context, attrName string, value string, owner sdk.AccAddress) {
	// Whole list for all the attributes with the same value

	attrs := k.FindAttributes(ctx, attrName, owner)
	store := ctx.KVStore(k.storeKey)

	if attrs == nil {
		panic("The property does not exists!")
	}
	// Update all the values
	for _, attr := range attrs {
		attr.Value = value // Change the value

		attrKey := k.AttributeStoreKey(owner, attrName)
		store.Set(attrKey, k.cdc.MustMarshalBinaryBare(attr))
	}
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetAttributesIterator(ctx sdk.Context, owner sdk.AccAddress) sdk.Iterator {
	print("GetAttributesIterator: owner=", owner)

	store := ctx.KVStore(k.storeKey)
	print("GetAttributesIterator: store=", store)

	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) FindAttributes(ctx sdk.Context, attrName string, owner sdk.AccAddress) []types.DemosAttribute {
	keyToFind := k.AttributeStoreKey(owner, attrName)
	iterator := k.GetAttributesIterator(ctx, owner)

	var result []types.DemosAttribute
	for ; iterator.Valid(); iterator.Next() {

		if bytes.Equal(keyToFind, iterator.Key()) {
			var attr types.DemosAttribute
			k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &attr)
			result = append(result, attr)
		}
	}

	return result
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string, owner sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	attrKey := k.AttributeStoreKey(owner, name)
	return store.Has(attrKey)
}

// Mixed key using the owner and the name
func (k Keeper) AttributeStoreKey(owner sdk.AccAddress, name string) []byte {
	return append(owner.Bytes(), []byte(name)...)
}
