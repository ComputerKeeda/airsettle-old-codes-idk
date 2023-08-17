package keeper

import (
	"context"

	"airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitValidatorVote(goCtx context.Context, msg *types.MsgSubmitValidatorVote) (*types.MsgSubmitValidatorVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitValidatorVoteResponse{}, nil
}
