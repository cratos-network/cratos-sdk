package keeper

import (
	"cratos.network/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the demosid Querier
const (
	QueryAll = "all"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {

		switch path[0] {
		case QueryAll:
			return queryAll(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown Demos ID query endpoint")
		}
	}
}

// nolint: unparam
func queryAll(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var result []types.DemosAttribute

	owner := sdk.AccAddress(req.GetData())
	iterator := keeper.GetAttributesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		//TODO: As in FindAttributes: a best approach could be to compare the key (bytes-wise) only to avoid to "open" each key
		attr := keeper.GetAttributeFromKey(ctx, iterator.Key())

		if attr.Owner.Equals(owner) {
			result = append(result, attr)
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
