package airsettle_test

import (
	"testing"

	keepertest "airsettle/testutil/keeper"
	"airsettle/testutil/nullify"
	"airsettle/x/airsettle"
	"airsettle/x/airsettle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AirsettleKeeper(t)
	airsettle.InitGenesis(ctx, *k, genesisState)
	got := airsettle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
