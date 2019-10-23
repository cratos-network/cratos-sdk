package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResResolve Queries Result Payload for a resolve query
type QueryResValue struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Namespace string `json:"namespace"`
}

// implement fmt.Stringer
func (res QueryResValue) String() string {
	return fmt.Sprintf("%s.%s %s", res.Namespace, res.Name, res.Value)
}

type QueryAllParams struct {
	Address sdk.AccAddress
}

type QueryValueParams struct {
	Address sdk.AccAddress `json:"address"`
	Value   string         `json:"value"`
}

// QueryResNames Queries Result Payload for a names query
type QueryResAll []DemosAttribute

// implement fmt.Stringer
func (list QueryResAll) String() string {
	var result string
	for i, attr := range list {
		result = result + fmt.Sprintf("% $.$=$", i, attr.Namespace, attr.Name, attr.Value)
	}

	return result
}
