package keeper

import (
	"bytes"
	"fmt"

	"cratos.network/cratos/x/demosid/internal/types"
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

// Gets the entire attribute metadata struct from a known key
func (k Keeper) GetAttributeFromKey(ctx sdk.Context, key []byte) types.DemosAttribute {

	store := ctx.KVStore(k.storeKey)
	storedBytes := store.Get(key)
	if storedBytes == nil {
		return types.DemosAttribute{}
	}

	var attr types.DemosAttribute
	k.cdc.MustUnmarshalBinaryBare(storedBytes, &attr)

	println(fmt.Sprint("--------------------------| %s=%s", attr.Name, attr.Value))

	return attr
}

// Gets the entire attribute metadata struct for a name
func (k Keeper) GetAttribute(ctx sdk.Context, namespace string, name string, owner sdk.AccAddress) types.DemosAttribute {

	nameKey := k.AttributeStoreKey(owner, namespace, name)
	return k.GetAttributeFromKey(ctx, nameKey)
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteAttribute(ctx sdk.Context, namespace string, attrName string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	nameKey := k.AttributeStoreKey(owner, namespace, attrName)
	store.Delete(nameKey)
}

// Changes the value for an attribute
func (k Keeper) SetValue(ctx sdk.Context, namespace string, attrName string, value string, owner sdk.AccAddress) {
	// Whole list for all the attributes with the same value
	store := ctx.KVStore(k.storeKey)
	attributeKey := k.AttributeStoreKey(owner, namespace, attrName)

	if k.IsNamePresent(ctx, namespace, attrName, owner) {
		// Update all the values
		println("Update the setting")
		attribute := k.GetAttribute(ctx, namespace, attrName, owner)
		// Update the value
		attribute.Value = value // Change the value
		store.Set(attributeKey, k.cdc.MustMarshalBinaryBare(attribute))
	} else {
		println("Create the setting")

		newAttribute := types.NewDemosAttribute(owner)
		// Setting the values
		newAttribute.Name = attrName
		newAttribute.Namespace = "Common"
		newAttribute.Value = value
		newAttribute.Owner = owner

		store.Set(attributeKey, k.cdc.MustMarshalBinaryBare(newAttribute))
	}
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetAttributesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func (k Keeper) FindAttributes(ctx sdk.Context, attrName string, owner sdk.AccAddress) []types.DemosAttribute {
	iterator := k.GetAttributesIterator(ctx)

	println("Reading attributes for ", owner)

	var result []types.DemosAttribute
	for ; iterator.Valid(); iterator.Next() {
		var attr types.DemosAttribute
		//TODO: The code open all the keys to compare it with the owner, namespace, ... The best approach could be to check the key
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &attr)

		if attr.Owner.Equals(owner) { // Only the keys for the owner
			result = append(result, attr)
		}
	}

	iterator.Close()
	return result
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, namespace string, name string, owner sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	attrKey := k.AttributeStoreKey(owner, namespace, name)
	return store.Has(attrKey)
}

// Mixed key using the owner and the name
func (k Keeper) AttributeStoreKey(owner sdk.AccAddress, namespace string, name string) []byte {
	// Join all the parts together
	result := bytes.Join([][]byte{owner.Bytes(), []byte(namespace), []byte(name)}, []byte{0x0})
	return result
}
