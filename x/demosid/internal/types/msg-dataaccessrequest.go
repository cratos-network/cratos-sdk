package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DataAccessRequestTypeName = "demosID/DataAccessRequest"
)

// MsgDataAccessRequest defines a AccessGrantRequest message
type MsgDataAccessRequest struct {
	From      sdk.AccAddress `json:"from"`
	DataOwner sdk.AccAddress `json:"owner"`
	Scope     string         `json:"scope"`
}

// Constructor function for MsgAccessGrantRequest
func NewMsgDataAccessRequest(fromAddr sdk.AccAddress, dataOwner sdk.AccAddress, scope string) MsgDataAccessRequest {
	return MsgDataAccessRequest{
		From:      fromAddr,
		DataOwner: dataOwner,
		Scope:     scope,
	}
}

// Route should return the name of the module
func (msg MsgDataAccessRequest) Route() string { return ModuleName }

// Type should return the action
func (msg MsgDataAccessRequest) Type() string { return DataAccessRequestTypeName }

// ValidateBasic runs stateless checks on the message
func (msg MsgDataAccessRequest) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}

	if msg.DataOwner.Empty() {
		return sdk.ErrInvalidAddress(msg.DataOwner.String())
	}

	if len(msg.Scope) == 0 {
		ErrInvalidScope()
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDataAccessRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDataAccessRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgDataAccessRequest) String() string {
	return fmt.Sprintf("%s ==> %s (%s)", msg.From, msg.DataOwner.Bytes(), msg.Scope)
}
