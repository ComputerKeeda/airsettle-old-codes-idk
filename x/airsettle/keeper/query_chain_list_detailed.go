package keeper

import (
	"context"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChainListDetailed(goCtx context.Context, req *types.QueryChainListDetailedRequest) (*types.QueryChainListDetailedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, err := k.GetAdminChainList(ctx, req.CreatorAddress)
	if !err {
		// data not found
		return &types.QueryChainListDetailedResponse{}, status.Error(codes.InvalidArgument, "No chains under this address")
	}

	var executionLayers []*types.Exelayer

	for i := 0; i < len(val.Id); i++ {
		value, found := k.GetExelayerById(ctx, val.Id[i])
		if !found {
			return &types.QueryChainListDetailedResponse{}, status.Error(codes.InvalidArgument, "Error in backend in getting details of chainid")
		}
		executionLayers = append(executionLayers, &value)
	}
	_ = executionLayers

	var pageResp *query.PageResponse

	return &types.QueryChainListDetailedResponse{
		Exelayer:   executionLayers,
		Pagination: pageResp,
	}, nil
}
