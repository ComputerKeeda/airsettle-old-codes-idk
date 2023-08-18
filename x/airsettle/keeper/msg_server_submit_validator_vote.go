package keeper

import (
	"context"
	"errors"
	"strconv"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitValidatorVote(goCtx context.Context, msg *types.MsgSubmitValidatorVote) (*types.MsgSubmitValidatorVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// type Poll struct {
	// 	PollId          string   `protobuf:"bytes,1,opt,name=pollId,proto3" json:"pollId,omitempty"`
	// 	ChainId         string   `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	// 	NewValidator    string   `protobuf:"bytes,3,opt,name=newValidator,proto3" json:"newValidator,omitempty"`
	// 	VotesDoneBy     []string `protobuf:"bytes,4,rep,name=votesDoneBy,proto3" json:"votesDoneBy,omitempty"`
	// 	Votes           []string `protobuf:"bytes,5,rep,name=votes,proto3" json:"votes,omitempty"`
	// 	TotalValidators uint64   `protobuf:"varint,6,opt,name=totalValidators,proto3" json:"totalValidators,omitempty"`
	// 	IsComplete      bool     `protobuf:"varint,7,opt,name=isComplete,proto3" json:"isComplete,omitempty"`
	// 	StartDate       string   `protobuf:"bytes,8,opt,name=startDate,proto3" json:"startDate,omitempty"`
	// 	PollCreator     string   `protobuf:"bytes,9,opt,name=pollCreator,proto3" json:"pollCreator,omitempty"`
	// }

	pollDetails, found := k.GetPollById(ctx, msg.PollId)
	if !found {
		Log("Cannot find poll details for poll id: " + msg.PollId)
		return nil, errors.New("Cannot find poll details for poll id: " + msg.PollId)
	}

	// checking if the creator || sender  have already voted or not

	var votesDoneByLength = len(pollDetails.VotesDoneBy)
	for i := 0; i < votesDoneByLength; i++ {
		votedBy := pollDetails.VotesDoneBy[i]
		if votedBy == msg.Creator {
			Log("Already voted")
			return &types.MsgSubmitValidatorVoteResponse{
				Success:     false,
				PollResult:  "--",
				Message:     "Already voted",
				Description: "Already voted",
			}, errors.New("Already voted")
		}
	}

	exeLayerDetails, found := k.GetExelayerById(ctx, pollDetails.ChainId)

	if !found {
		Log("Execution layer not found")
		return &types.MsgSubmitValidatorVoteResponse{
			Success:     false,
			PollResult:  "--",
			Message:     "Execution layer not found",
			Description: "Execution layer not found",
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
		return &types.MsgSubmitValidatorVoteResponse{
			Success:     false,
			PollResult:  "--",
			Message:     "Requester is not a validator",
			Description: "Requester is not a validator",
		}, errors.New("Requester is not a validator")
	}

	var newVotesDoneBy = pollDetails.VotesDoneBy
	newVotesDoneBy = append(newVotesDoneBy, msg.Creator)

	var newVotes = pollDetails.Votes
	myStringVoteValue := strconv.FormatBool(msg.Vote)
	newVotes = append(newVotes, myStringVoteValue)

	// checking if the poll is complete or not
	var newIsComplete bool
	if len(newVotesDoneBy) >= int(pollDetails.TotalValidators) {
		// ! poll is complete
		newIsComplete = true
		// And if the poll is complete then we have to check the votes

		// ! checking the votes
		var successVotePercentage float64 = 00.0
		var trueVoteCount int = 0
		for i := 0; i < len(pollDetails.Votes); i++ {
			value := pollDetails.Votes[i]
			if value == "true" {
				trueVoteCount++
			}
		}

		successVotePercentage = float64(trueVoteCount) / float64(pollDetails.TotalValidators) * 100

		if successVotePercentage >= 50.0 {
			var newValidators = exeLayerDetails.Validator
			var newVotingPower = exeLayerDetails.VotingPower
			newValidators = append(newValidators, pollDetails.NewValidator)
			newVotingPower = append(newVotingPower, 100)

			k.UpdateExecutionLayers(ctx, types.Exelayer{
				Validator:            newValidators,
				VotingPower:          newVotingPower,
				LatestBatch:          exeLayerDetails.LatestBatch,
				LatestMerkleRootHash: exeLayerDetails.LatestMerkleRootHash,
				VerificationKey:      exeLayerDetails.VerificationKey,
				ChainInfo:            exeLayerDetails.ChainInfo,
				Id:                   exeLayerDetails.Id,
				Creator:              exeLayerDetails.Creator,
			})
		}
	} else {
		newIsComplete = false
	}

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix))
	var poll = types.Poll{
		PollId:          pollDetails.PollId,
		ChainId:         pollDetails.ChainId,
		NewValidator:    pollDetails.NewValidator,
		VotesDoneBy:     newVotesDoneBy,
		Votes:           newVotes,
		TotalValidators: pollDetails.TotalValidators,
		IsComplete:      newIsComplete,
		StartDate:       pollDetails.StartDate,
		PollCreator:     pollDetails.PollCreator,
	}

	b := k.cdc.MustMarshal(&poll)
	pollStore.Set([]byte(pollDetails.PollId), b)

	// ?ignore this for now
	return &types.MsgSubmitValidatorVoteResponse{
		Success:     true,
		PollResult:  "some result will come here",
		Message:     "Your vote has been submitted successfully",
		Description: "some description will come here",
	}, nil
}
