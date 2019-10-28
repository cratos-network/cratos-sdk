package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidScope sdk.CodeType = 1001
)

// ErrNameDoesNotExist is the error for name not existing
func ErrInvalidScope() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidScope, "Invalid scope")
}
