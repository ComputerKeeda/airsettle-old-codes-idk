package keeper

import (
	"airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetPollById(ctx sdk.Context, PollIDFromReq string) (poll types.Poll, found bool) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKeyPrefix))

	store := ctx.KVStore(k.storeKey)
	poll_Store := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix))

	b := poll_Store.Get([]byte(PollIDFromReq))

	var pollDetails types.Poll
	if b == nil {
		Log("Cannot find poll details for poll id: " + PollIDFromReq)
		return pollDetails, false
	}

	k.cdc.MustUnmarshal(b, &pollDetails)
	return pollDetails, true
}
