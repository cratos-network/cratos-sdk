package demosid

import (
	"aquarelle.io/cratos/x/demosid/internal/keeper"
	"aquarelle.io/cratos/x/demosid/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper             = keeper.NewKeeper
	NewQuerier            = keeper.NewQuerier
	NewMsgDeleteAttribute = types.NewMsgDeleteAttribute
	NewMsgSetAttribute    = types.NewMsgSetAttribute
	NewDemosAttribute     = types.NewDemosAttribute
	ModuleCdc             = types.ModuleCdc
	RegisterCodec         = types.RegisterCodec
)

type (
	Keeper             = keeper.Keeper
	MsgSetAttribute    = types.MsgSetAttribute
	MsgDeleteAttribute = types.MsgDeleteAttribute
	QueryResAll        = types.QueryResAll
)
