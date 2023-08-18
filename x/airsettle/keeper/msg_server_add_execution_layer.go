package keeper

import (
	"context"

	"airsettle/x/airsettle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) AddExecutionLayer(goCtx context.Context, msg *types.MsgAddExecutionLayer) (*types.MsgAddExecutionLayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newUUID := uuid.New().String()

	var exelayer = types.Exelayer{
		Validator:            []string{msg.Creator},
		VotingPower:          []uint64{100},
		LatestBatch:          0,
		LatestMerkleRootHash: "0",
		VerificationKey:      msg.VerificationKey,
		ChainInfo:            msg.ChainInfo,
		Id:                   newUUID,
		Creator:              msg.Creator,
	}

	k.SetExecutionLayers(
		ctx,
		exelayer,
	)

	LogCreateFileOnPath(newUUID, "test/chainid.test.air")
	Log("Execution layer created. chainId: " + newUUID)
	return &types.MsgAddExecutionLayerResponse{Id: newUUID}, nil
}
