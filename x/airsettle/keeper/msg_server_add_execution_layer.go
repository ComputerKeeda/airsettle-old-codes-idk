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

	k.SetExecutionlayers(
		ctx,
		exelayer,
	)

	LogCreateChainid(newUUID)
	Log("Execution layer created. chainId: " + newUUID)
	return &types.MsgAddExecutionLayerResponse{Id: newUUID}, nil
}

// func WritingValuesToFile(newUUID, creatorAddress string) {
// 	// Get the directory of the testing.sh script
// 	// update the path below
// 	scriptPath := "test/testing.sh"
// 	scriptDir := filepath.Dir(scriptPath)

// 	// Determine the path to the updated_values.txt file
// 	valuesFilePath := filepath.Join(scriptDir, "updated_values.txt")

// 	newContent := fmt.Sprintf("chainid=\"%s\"\ncreator_address=\"%s\"\n", newUUID, creatorAddress)

// 	err := os.WriteFile(valuesFilePath, []byte(newContent), 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
