package cli

import (
	"github.com/spf13/cobra"

	"cratos.network/cratos/x/demosid/internal/types"
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
		GetCmdDataAccessGrantRequest(cdc),
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
func GetCmdDataAccessGrantRequest(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request [namespace] [name]",
		Short: "Request for access to the non-public attributes of other account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			ownerAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			scope := args[1]

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDataAccessRequest(cliCtx.GetFromAddress(), ownerAddr, scope)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
