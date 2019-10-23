package demosid

import (
	"fmt"

	"aquarelle.io/cratos/x/demosid/internal/types"
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

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Attributes {
		keeper.SetValue(ctx, record.Name, record.Value, record.Owner)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []types.DemosAttribute

	accounts := k.AccountKeeper.GetAllAccounts(ctx)

	for _, account := range accounts {
		owner := account.GetAddress()
		iterator := k.GetAttributesIterator(ctx, owner)
		for ; iterator.Valid(); iterator.Next() {

			name := string(iterator.Key())
			attr := k.GetAttribute(ctx, name, owner)
			records = append(records, attr)

		}
	}
	return GenesisState{Attributes: records}
}
