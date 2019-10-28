package keeper

import (
	"bytes"

	"cratos.network/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the demosid Querier
const (
	QueryAll           = "all"
	QuerierAllRequests = "access"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {

		switch path[0] {
		case QueryAll:
			return queryAll(ctx, req, keeper)
		case QuerierAllRequests:
			return queryAllRequests(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown Demos ID query endpoint")
		}
	}
}

// nolint: unparam
func queryAll(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var result []types.DemosAttribute

	payload := bytes.Split(req.GetData(), types.DefaultKeySeparator)
	owner := sdk.AccAddress(payload[0])
	fromAddr := sdk.AccAddress(payload[1])

	// Only when the requester is the same user, the operation is fully granted. Otherwise, the code will check the request
	isGranted := owner.Equals(fromAddr)

	iterator := keeper.GetAttributesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if key[0] == types.AttributeBytesHeader[0] { // Only attributes
			if isGranted { // Only show the data to the same account
				attr := keeper.GetAttributeFromKey(ctx, key)

				if attr.Owner.Equals(owner) {
					result = append(result, attr)
				}
			}
		}
	}

	if len(result) == 0 {
		return []byte{}, sdk.ErrInvalidSequence("There are not any attributes defined")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryAllRequests(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var result []types.DataAccessRequest

	// owner := sdk.AccAddress(req.GetData())
	iterator := keeper.GetAttributesIterator(ctx)
	store := ctx.KVStore(keeper.storeKey)

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		if key[0] == types.DataAccessGrantRequestBytesHeader[0] { // Only requests
			var request types.DataAccessRequest
			keeper.cdc.MustUnmarshalBinaryBare(store.Get(key), &request)

			result = append(result, request)
		}
	}

	if len(result) == 0 {
		return []byte{}, nil
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
