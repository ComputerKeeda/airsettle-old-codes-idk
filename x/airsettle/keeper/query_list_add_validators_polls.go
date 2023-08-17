package keeper

import (
	"context"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAddValidatorsPolls(goCtx context.Context, req *types.QueryListAddValidatorsPollsRequest) (*types.QueryListAddValidatorsPollsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKeyPrefix))

	var pollIds []string
	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error { // Make sure req.Pagination is defined
		var poll types.Poll
		if err := k.cdc.Unmarshal(value, &poll); err != nil {
			return err
		}
		pollIds = append(pollIds, poll.PollId) // Fixing the field name from pollId to PollId
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListAddValidatorsPollsResponse{PollIds: pollIds}, nil
}
