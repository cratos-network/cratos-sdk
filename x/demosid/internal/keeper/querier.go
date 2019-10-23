package keeper

import (
	"fmt"

	"aquarelle.io/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the demosid Querier
const (
	QueryAll   = "all"
	QueryValue = "value"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {

		switch path[0] {
		case QueryAll:
			return queryAll(ctx, req, keeper)
		case QueryValue:
			return queryValue(ctx, path[1], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown Demos ID query endpoint")
		}
	}
}

// nolint: unparam
func queryAll(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	address := sdk.AccAddress(req.GetData())
	iterator := keeper.GetAttributesIterator(ctx, address)

	var result []types.DemosAttribute
	for ; iterator.Valid(); iterator.Next() {
		attr := keeper.GetAttribute(ctx, string(iterator.Key()), address)
		result = append(result, attr)
	}

	if len(result) == 0 {
		return []byte{}, sdk.ErrInvalidSequence("The holder has not any attributes defined")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryValue(ctx sdk.Context, name string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.QueryAllParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := keeper.GetAttribute(ctx, name, params.Address)

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
