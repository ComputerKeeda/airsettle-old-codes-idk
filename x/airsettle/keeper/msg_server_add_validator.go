package keeper

import (
	"context"
	"errors"
	"strings"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func isValidCosmosAddress(address string) bool {
	const customPrefix = "air"
	// Check if the address has the correct prefix
	if !strings.HasPrefix(address, customPrefix) {
		return false
	}

	// Decode the Bech32 encoded address
	_, _, err := bech32.Decode(address, bech32.MaxLengthBIP173)

	// Return true if decoding was successful, false otherwise
	return err == nil
}

func (k msgServer) AddValidator(goCtx context.Context, msg *types.MsgAddValidator) (*types.MsgAddValidatorResponse, error) {

	Log("AddValidator Called with ChainId: " + msg.ChainId + " and NewValidatorAddress: " + msg.NewValidatorAddress)
	ctx := sdk.UnwrapSDKContext(goCtx)
	newUUID := uuid.New().String()
	// Validate the NewValidatorAddress Address

	if !isValidCosmosAddress(msg.NewValidatorAddress) {
		Log("Invalid NewValidatorAddress")
		return &types.MsgAddValidatorResponse{
			VotingPollId: "--",
		}, errors.New("Invalid NewValidatorAddress")
	}

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

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix))

	// check if the NewValidatorAddress is already in the list or not

	iterator := sdk.KVStorePrefixIterator(pollStore, []byte{})
	var polls []types.Poll
	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshal(iterator.Value(), &poll)
		polls = append(polls, poll)
		Log("PollId: " + poll.PollId + " NewValidatorAddress: " + poll.NewValidator + " ChainId: " + poll.ChainId)
	}

	iterator.Close()

	for i := 0; i < len(polls); i++ {
		poll := polls[i]
		if poll.NewValidator == msg.NewValidatorAddress && poll.ChainId == msg.ChainId {
			Log("NewValidatorAddress is already in the list at PollId: " + poll.PollId)
			return &types.MsgAddValidatorResponse{
				VotingPollId: poll.PollId,
			}, errors.New("NewValidatorAddress is already in the list at PollId: " + poll.PollId)
		}
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

	b := k.cdc.MustMarshal(&poll)
	pollStore.Set([]byte(newUUID), b)

	LogLoop([]string{"UUID created", newUUID})
	LogCreateFileOnPath(newUUID, "test/pollid.test.air")
	return &types.MsgAddValidatorResponse{
		VotingPollId: newUUID,
	}, nil
}
