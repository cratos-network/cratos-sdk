package demosid

import (
	"fmt"

	"cratos.network/cratos/x/demosid/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Attributes []types.DemosAttribute `json:"attributes"`
}

func NewGenesisState(attributes []types.DemosAttribute) GenesisState {
	return GenesisState{Attributes: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Attributes {
		if record.Owner == nil {
			return fmt.Errorf("Invalid Attribute: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Namespace == "" {
			return fmt.Errorf("Invalid Attribute: Owner: %s. Error: Missing Namespace", record.Owner)
		}
		if record.Name == "" {
			return fmt.Errorf("Invalid Attribute: Value: %s. Error: Missing name of the property", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Attributes: []types.DemosAttribute{},
	}
}

// Initialize the genesis re-creating each attribute
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Attributes {
		keeper.SetValue(ctx, record.Namespace, record.Name, record.Value, record.Owner)
	}
	return []abci.ValidatorUpdate{}
}

// Export all the attributes
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []types.DemosAttribute

	iterator := k.GetAttributesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		//Read the attribute
		attr := k.GetAttributeFromKey(ctx, iterator.Key())
		records = append(records, attr)
	}

	return GenesisState{Attributes: records}
}
