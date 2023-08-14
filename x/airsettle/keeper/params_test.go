package keeper_test

import (
	"testing"

	testkeeper "airsettle/testutil/keeper"
	"airsettle/x/airsettle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AirsettleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
