package cli

import (
	"github.com/spf13/cobra"

	"aquarelle.io/cratos/x/demosid/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	demosIDTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Cratos ID transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	demosIDTxCmd.AddCommand(client.PostCommands(
		GetCmdSetAttribute(cdc),
		GetCmdDeleteAttribute(cdc),
	)...)

	return demosIDTxCmd
}

// GetCmdBuyName is the CLI command for sending a BuyName transaction
func GetCmdSetAttribute(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <name> <value>",
		Short: "Set a value for existing attribute or create a new one",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			attributeName := args[0]
			attributeValue := args[1]
			key := cliCtx.GetFromAddress()

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgSetAttribute(attributeName, attributeValue, "common", key)
			err := msg.ValidateBasic()
			if err != nil {
				println("Error validating the message")
				return err
			}

			result := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			return result
		},
	}

	return cmd
}

// GetCmdSetName is the CLI command for sending a SetName transaction
func GetCmdDeleteAttribute(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete am attribute",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteAttribute(args[0], "common", cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
