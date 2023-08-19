package keeper

import (
	"context"

	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"

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

	pollDetails, found := k.GetPollById(ctx, PollIDFromReq)

	if !found {
		Log("Cannot find poll details for poll id: " + PollIDFromReq)
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	return &types.QueryValidatorPollDetailsResponse{Poll: &pollDetails}, nil
}
