package keeper

import (
	"github.com/airchains-network/Airchains-settlement-layer/x/airsettle/types"
)

var _ types.QueryServer = Keeper{}
