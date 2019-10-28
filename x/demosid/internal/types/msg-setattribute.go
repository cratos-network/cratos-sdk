package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	SetAttributeTypeName = "demosID/SetAttribute"
)

// MsgSetName defines a SetName message
type MsgSetAttribute struct {
	Name      string         `json:"name"`
	Value     string         `json:"value"`
	Namespace string         `json:"namespace"`
	Owner     sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetAttribute(name string, value string, namespace string, owner sdk.AccAddress) MsgSetAttribute {
	return MsgSetAttribute{
		Name:      name,
		Value:     value,
		Namespace: namespace,
		Owner:     owner,
	}
}

// Route should return the name of the module
func (msg MsgSetAttribute) Route() string { return ModuleName }

// Type should return the action
func (msg MsgSetAttribute) Type() string { return SetAttributeTypeName }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetAttribute) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 || len(msg.Value) == 0 || len(msg.Namespace) == 0 {
		return sdk.ErrUnknownRequest("Name, Namespace and/or Value cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetAttribute) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetAttribute) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgSetAttribute) String() string {
	return fmt.Sprintf("%s ==> %s (%s) = %s", msg.Owner, msg.Name, msg.Namespace, msg.Value)
}
