package cli

import (
	"strconv"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdListAddValidatorsPolls() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-add-validators-polls",
		Short: "Query list_add_validators_polls",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListAddValidatorsPollsRequest{}

			res, err := queryClient.ListAddValidatorsPolls(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
