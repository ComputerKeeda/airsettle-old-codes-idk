package keeper

import (
	"context"
	"errors"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) AddValidator(goCtx context.Context, msg *types.MsgAddValidator) (*types.MsgAddValidatorResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	newUUID := uuid.New().String()

	exeLayerDetails, found := k.GetExelayerById(ctx, msg.ChainId)

	if !found {
		Log("Execution layer not found")
		return &types.MsgAddValidatorResponse{
			VotingPollId: "--",
		}, errors.New("Execution layer not found")
	}

	var validatorsLength = len(exeLayerDetails.Validator)

	//* checking if sender is a validator or not
	var isAuthenticValidator = false
	for i := 0; i < validatorsLength; i++ {
		validatorAddress := exeLayerDetails.Validator[i] // ? already present validator address
		if validatorAddress == msg.Creator {
			isAuthenticValidator = true
			break
		}
	}

	if !isAuthenticValidator {
		Log("Requester is not a validator")
		return &types.MsgAddValidatorResponse{
			VotingPollId: "--",
		}, errors.New("Requester is not a validator")
	}

	var poll = types.Poll{
		PollId:          newUUID,
		ChainId:         msg.ChainId,
		NewValidator:    msg.NewValidatorAddress,
		VotesDoneBy:     []string{msg.Creator},
		Votes:           []string{"true"},
		TotalValidators: uint64(validatorsLength),
		IsComplete:      false,
		StartDate:       ctx.BlockTime().String(),
		PollCreator:     msg.Creator,
	}

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix))
	b := k.cdc.MustMarshal(&poll)
	pollStore.Set([]byte(newUUID), b)

	LogLoop([]string{"UUID created", newUUID})

	return &types.MsgAddValidatorResponse{
		VotingPollId: newUUID,
	}, nil
}
