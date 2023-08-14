package cli

import (
	"strconv"

	"airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChainList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-list [creator-address]",
		Short: "Query chain_list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCreatorAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryChainListRequest{

				CreatorAddress: reqCreatorAddress,
			}

			res, err := queryClient.ChainList(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
