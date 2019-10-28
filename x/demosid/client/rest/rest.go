package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "attribute"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/all", storeName), allHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s", storeName), valueHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s", storeName), setAttributeHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/access/%s/%s", storeName), requestDataAccessHandler(cliCtx)).Methods("POST")
}
