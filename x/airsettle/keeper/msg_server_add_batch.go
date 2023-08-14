package keeper

import (
	"airsettle/x/airsettle/types"
	"context"
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBatch(goCtx context.Context, msg *types.MsgAddBatch) (*types.MsgAddBatchResponse, error) {
	// ! Check Here is the main error (AT ASS)
	ctx := sdk.UnwrapSDKContext(goCtx)
	// todo: return error & message & bool not just bool

	Log("Creating Batch. chainId:" + msg.Id + ", Batch Number:" + strconv.Itoa(int(msg.BatchNumber)))

	// check if execution layer aka chainid exists
	exelayer, found := k.GetExelayerById(ctx, msg.Id)
	if !found {
		Log("Execution layer dont exists")
		return &types.MsgAddBatchResponse{
			BatchStatus: false, // execution layer don't exists
		}, nil
	}

	// check if batch number is correct
	if exelayer.LatestBatch+1 != msg.BatchNumber {
		Log("Wrong batch number")
		return &types.MsgAddBatchResponse{
			BatchStatus: false, // Wrong batch number
		}, nil
	}

	// check if msg.Sender is a Validator/Admin
	isValidator := false
	for _, value := range exelayer.Validator {
		if value == msg.Creator {
			isValidator = true
		}
	}

	if !isValidator {
		Log("Not A Validator, Batch not created")
		return &types.MsgAddBatchResponse{
			BatchStatus: false, // NOT a validator
		}, nil
	}

	// todo: optional: (not required, not good for testing) check if prevMerkleRootHash of sender is equal to current LatestMekleRootHash

	// Add Batch on chain id
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix("/Batch/"+msg.Id+"/"))
	appendedValue := k.cdc.MustMarshal(msg)
	store.Set(GetBytesFromUint64(msg.BatchNumber), appendedValue)

	// update Executionlayer data: latest batch, merkle roots
	exelayer.LatestBatch = msg.BatchNumber
	exelayer.LatestMerkleRootHash = msg.MerkleRootHash
	exelayer.VerificationKey = "/getvkey/" + exelayer.Id + "/" // not changed
	k.UpdateExecutionlayers(ctx, exelayer)

	Log("Batch " + strconv.Itoa(int(msg.BatchNumber)) + " Created")
	return &types.MsgAddBatchResponse{
		BatchStatus: true, // success
	}, nil
}

func GetBytesFromUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
