package keeper

import (
	"context"

	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowExecutionLayer(goCtx context.Context, req *types.QueryShowExecutionLayerRequest) (*types.QueryShowExecutionLayerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	exelayer, found := k.GetExelayerById(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	vkey := k.LocalGetvkey(ctx, req.Id)

	exelayer_data := types.Exelayer{
		Validator:            exelayer.Validator,
		VotingPower:          exelayer.VotingPower,
		LatestBatch:          exelayer.LatestBatch,
		LatestMerkleRootHash: exelayer.LatestMerkleRootHash,
		ChainInfo:            exelayer.ChainInfo,
		VerificationKey:      vkey,
		Id:                   exelayer.Id,
		Creator:              exelayer.Creator,
	}
	LogLoop(exelayer_data.Validator)
	return &types.QueryShowExecutionLayerResponse{Exelayer: exelayer_data}, nil
}
