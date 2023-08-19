package airsettle_test

import (
	"testing"

	keepertest "github.com/Airchains-Studio/Settlement_Layer/testutil/keeper"
	"github.com/Airchains-Studio/Settlement_Layer/testutil/nullify"
	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle"
	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"
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
