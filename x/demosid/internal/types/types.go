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
	IsPublic  bool           `json:"isPublic"`
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
	return strings.TrimSpace(fmt.Sprintf(`
Name: %s
Value: %s
Namespace: %s
Owner: %s
`, dattr.Name, dattr.Value, dattr.Namespace, dattr.Owner))
}

type DataAccessRequest struct {
	From      sdk.AccAddress `json:"from"`
	DataOwner sdk.AccAddress `json:"owner"`
	Scope     string         `json:"scope"`
	OpKey     string         `json:""opkey"`
}

func NewDataAccessRequest(fromAddr sdk.AccAddress, dataOwner sdk.AccAddress, scope string) DataAccessRequest {
	return DataAccessRequest{
		From:      fromAddr,
		DataOwner: dataOwner,
		Scope:     scope,
	}
}

func (request DataAccessRequest) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
From: %s
Data owner: %s
Scope list: %s
`, request.From, request.DataOwner, request.Scope))
}
