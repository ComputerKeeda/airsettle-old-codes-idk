package cli

import (
	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdAddExecutionLayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-execution-layer [chain-info] [verification-key-json-file]",
		Short: "Broadcast message add_execution_layer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			verificationKeyFilePath := args[1]

			verificationKeyJSON, err := readJSON(verificationKeyFilePath)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddExecutionLayer(
				clientCtx.GetFromAddress().String(),
				string(verificationKeyJSON),
				args[0],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
