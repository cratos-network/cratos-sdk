package keeper

import (
	"bytes"
	"strings"

	"cratos.network/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/google/uuid"
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

	println("Primer byte = ", key[0])

	store := ctx.KVStore(k.storeKey)
	storedBytes := store.Get(key)
	if storedBytes == nil {
		return types.DemosAttribute{}
	}

	var attribute types.DemosAttribute
	k.cdc.MustUnmarshalBinaryBare(storedBytes, &attribute)
	return attribute
}

// Gets the entire attribute metadata struct for a name
func (k Keeper) GetAttribute(ctx sdk.Context, namespace string, name string, owner sdk.AccAddress) types.DemosAttribute {

	nameKey := k.GetAttributeStoreKey(owner, namespace, name)
	return k.GetAttributeFromKey(ctx, nameKey)
}

// Changes the value for an attribute
func (k Keeper) SetValue(ctx sdk.Context, namespace string, attrName string, value string, owner sdk.AccAddress) {
	// Whole list for all the attributes with the same value
	store := ctx.KVStore(k.storeKey)
	key := k.GetAttributeStoreKey(owner, namespace, attrName)

	if k.IsNamePresent(ctx, namespace, attrName, owner) {
		// Update all the values
		attribute := k.GetAttribute(ctx, namespace, attrName, owner)
		// Update the value
		attribute.Value = value // Change the value
		store.Set(key, k.cdc.MustMarshalBinaryBare(attribute))
	} else {

		newAttribute := types.NewDemosAttribute(owner)
		// Setting the values
		newAttribute.Name = attrName
		newAttribute.Namespace = namespace
		newAttribute.Value = value
		newAttribute.Owner = owner

		store.Set(key, k.cdc.MustMarshalBinaryBare(newAttribute))
	}
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetAttributesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func (k Keeper) CreateDataAccessRequest(ctx sdk.Context, fromAddr sdk.AccAddress, dataOwner sdk.AccAddress, scope string) {

	store := ctx.KVStore(k.storeKey)
	// split the scopes and create a request for each scope
	scopeList := strings.Split(scope, ",")

	for i := 0; i < len(scopeList); i++ {
		singleScope := scopeList[i]
		key := k.GetRequestStoreKey(fromAddr, dataOwner, singleScope)
		// Only create a new request if it does not exists previously
		if (store.Get(key)) == nil {
			request := types.NewDataAccessRequest(fromAddr, dataOwner, singleScope)
			// Creates a unique operation key (opkey) using the store key as data seed
			request.OpKey = uuid.NewSHA1(uuid.NameSpaceOID, key).String()

			store.Set(key, k.cdc.MustMarshalBinaryBare(request))
		}
	}
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, namespace string, name string, owner sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	attrKey := k.GetAttributeStoreKey(owner, namespace, name)
	return store.Has(attrKey)
}

// Mixed key using the owner and the name
func (k Keeper) GetAttributeStoreKey(owner sdk.AccAddress, namespace string, name string) []byte {
	// Join all the parts together
	return bytes.Join([][]byte{
		types.AttributeBytesHeader, // An initial header
		owner.Bytes(),
		[]byte(namespace),
		[]byte(name)}, []byte{0x0})
}

// Builds a key to store a data access request
func (k Keeper) GetRequestStoreKey(fromAddr sdk.AccAddress, dataOwner sdk.AccAddress, scope string) []byte {

	return bytes.Join([][]byte{
		types.DataAccessGrantRequestBytesHeader, // An initial header
		fromAddr.Bytes(),
		dataOwner.Bytes(),
		[]byte(scope)}, []byte{0x0})
}
