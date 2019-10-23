package demosid

import (
	"fmt"

	"cratos.network/cratos/x/demosid/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for Demos ID type messages.
func NewHandler(keeper Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		switch msg := msg.(type) {
		case types.MsgSetAttribute:
			return handleMsgSetAttribute(ctx, keeper, msg)
		case types.MsgDeleteAttribute:
			return handleMsgDeleteAttribute(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Cratos message type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetAttribute(ctx sdk.Context, keeper Keeper, msg MsgSetAttribute) sdk.Result {

	keeper.SetValue(ctx, msg.Namespace, msg.Name, msg.Value, msg.Owner) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                                                 // return
}

// Handle a message to delete name
func handleMsgDeleteAttribute(ctx sdk.Context, keeper Keeper, msg MsgDeleteAttribute) sdk.Result {

	keeper.DeleteAttribute(ctx, msg.Namespace, msg.Name, msg.Owner)
	return sdk.Result{}
}
