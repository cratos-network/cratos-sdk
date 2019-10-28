package demosid

import (
	"cratos.network/cratos/x/demosid/internal/keeper"
	"cratos.network/cratos/x/demosid/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper               = keeper.NewKeeper
	NewQuerier              = keeper.NewQuerier
	NewMsgSetAttribute      = types.NewMsgSetAttribute
	NewMsgDataAccessRequest = types.NewMsgDataAccessRequest
	NewDemosAttribute       = types.NewDemosAttribute
	ModuleCdc               = types.ModuleCdc
	RegisterCodec           = types.RegisterCodec
)

type (
	Keeper                   = keeper.Keeper
	MsgSetAttribute          = types.MsgSetAttribute
	MsgDataAccessRequest     = types.MsgDataAccessRequest
	QueryResGetAllAttributes = types.QueryResGetAllAttributes
	QueryResGetAllRequests   = types.QueryResGetAllRequests
)
