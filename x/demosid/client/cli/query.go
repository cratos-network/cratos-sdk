package cli

import (
	"fmt"

	"aquarelle.io/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	demosIDQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the Cratos ID module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	demosIDQueryCmd.AddCommand(client.GetCommands(
		GetCmdAllAttributes(storeKey, cdc),
		GetValue(storeKey, cdc),
	)...)

	return demosIDQueryCmd
}

// GetCmdAllAttributes queries the list of attributes
func GetCmdAllAttributes(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <address>",
		Short: "List all settings and their values",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all", queryRoute), address)
			if err != nil {
				fmt.Printf("could not list attributes: \nReason: %s \n", err)
				return nil
			}

			var out types.QueryResAll
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

	return cmd

}

// GetValue queries the value for an attribute
func GetValue(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "value [name] [address]",
		Short: "Query value of a attribute",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			accGetter := auth.NewAccountRetriever(cliCtx)

			name := args[0]
			key, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			if err := accGetter.EnsureExists(key); err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/value/%s", queryRoute, name), key)
			if err != nil {
				fmt.Printf("could not get the value for - %s \n", name)
				return nil
			}

			var out types.QueryResValue
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
