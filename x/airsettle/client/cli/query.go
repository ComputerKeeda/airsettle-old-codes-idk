package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"airsettle/x/airsettle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group airsettle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowExecutionLayer())

	cmd.AddCommand(CmdListExecutionLayers())

	cmd.AddCommand(CmdShowBatch())

	cmd.AddCommand(CmdChainList())

	cmd.AddCommand(CmdChainListDetailed())

	cmd.AddCommand(CmdVerificationKey())

	cmd.AddCommand(CmdVerify())

	cmd.AddCommand(CmdListAddValidatorsPolls())

	cmd.AddCommand(CmdValidatorPollDetails())

	// this line is used by starport scaffolding # 1

	return cmd
}
