package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CommonNamespaceName = "Common"
)

// MinIDPrice is Initial Starting Price for a ID that was never previously owned
var MinAttributePrice = sdk.Coins{sdk.NewInt64Coin("cratos", 1)}

// DemosID is a struct that contains all the metadata of a ID
type DemosAttribute struct {
	Name      string         `json:"name"`
	Value     string         `json:"value"`
	Namespace string         `json:"namespace"`
	Owner     sdk.AccAddress `json:"owner"`
}

// NewDemosID returns a new ID with the global namespace as default
func NewDemosAttribute(owner sdk.AccAddress) DemosAttribute {
	return DemosAttribute{
		Namespace: CommonNamespaceName,
	}
}

// implement fmt.Stringer
func (dattr DemosAttribute) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s
Value: %s
Namespace: %s
Owner: %s`, dattr.Name, dattr.Value, dattr.Namespace, dattr.Owner))
}
