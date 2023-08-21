package keeper

import (
	"github.com/airchains-network/Airchains-settlement-layer/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) LocalGetvkey(ctx sdk.Context, chainid string) string {
	// get verification key
	var vk_val types.Vkey
	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	b := vk_store.Get([]byte(chainid))
	if b == nil {
		vk_val.VerificationKey = ""
		// verification key not found
	} else {
		k.cdc.MustUnmarshal(b, &vk_val)
	}
	return vk_val.VerificationKey
}
