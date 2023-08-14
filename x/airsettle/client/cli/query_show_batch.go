package cli

import (
	"strconv"

	"airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdShowBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-batch [id] [batch-number]",
		Short: "Query show_batch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]
			reqBatchNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowBatchRequest{

				Id:          reqId,
				BatchNumber: reqBatchNumber,
			}

			res, err := queryClient.ShowBatch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
