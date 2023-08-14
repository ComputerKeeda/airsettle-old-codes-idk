package cli

import (
	"airsettle/x/airsettle/types"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func readJSON(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}
	return data, nil
}

func CmdAddBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-batch [id] [batch-number] [merkle-root-hash] [prev-merkle-root-hash] [zk-proof-json-file]",
		Short: "Broadcast message add_batch",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			jsonFilePath := args[4]

			jsonData, err := readJSON(jsonFilePath)
			if err != nil {
				return fmt.Errorf("failed to read JSON: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("failed to get client context: %w", err)
			}

			msg := types.NewMsgAddBatch(
				clientCtx.GetFromAddress().String(),
				args[0],
				cast.ToUint64(args[1]),
				args[2],
				args[3],
				string(jsonData),
			)
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
