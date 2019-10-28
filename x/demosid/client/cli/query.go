package cli

import (
	"bytes"
	"fmt"

	"cratos.network/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		GetCmdAllAccessRequests(storeKey, cdc),
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
				return nil
			}

			if cliCtx.FromAddress == nil || cliCtx.FromAddress.Empty() {
				return sdk.ErrInvalidAddress("The account requester must be specified. Use the flag --from")
			}

			// Send the request including the from
			payload := bytes.Join([][]byte{address, cliCtx.FromAddress}, types.DefaultKeySeparator)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all/%s", queryRoute), payload)
			if err != nil {
				return nil
			}

			var out types.QueryResGetAllAttributes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
	cmd.Flags().String(flags.FlagFrom, "", "Address of the private key of the Account who made the request")
	return cmd

}

func GetCmdAllAccessRequests(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access <address>",
		Short: "List all request for data access",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			ownerAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return nil
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/access", queryRoute), ownerAddr)
			if err != nil {
				return nil
			}

			var out types.QueryResGetAllRequests
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}

	return cmd
}
