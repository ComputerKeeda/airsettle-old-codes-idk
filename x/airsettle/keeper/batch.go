package keeper

import (
	"airsettle/x/airsettle/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetbatchById(ctx sdk.Context, id string, batchnumber uint64) (val types.MsgAddBatch, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix("/Batch/"+id+"/")) //types.PostKey)) // types.PostKey = "Post/value/"

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, batchnumber)

	b := store.Get(bz)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
