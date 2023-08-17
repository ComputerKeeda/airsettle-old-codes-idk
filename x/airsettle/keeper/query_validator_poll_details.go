package keeper

import (
	"context"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorPollDetails(goCtx context.Context, req *types.QueryValidatorPollDetailsRequest) (*types.QueryValidatorPollDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var PollIDFromReq = req.PollId

	store := ctx.KVStore(k.storeKey)
	poll_Store := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix))

	b := poll_Store.Get([]byte(PollIDFromReq))

	if b == nil {
		Log("Cannot find poll details for poll id: " + PollIDFromReq)
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pollDetails types.Poll

	k.cdc.MustUnmarshal(b, &pollDetails)

	return &types.QueryValidatorPollDetailsResponse{Poll: &pollDetails}, nil
}
