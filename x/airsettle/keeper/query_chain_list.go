package keeper

import (
	"context"

	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChainList(goCtx context.Context, req *types.QueryChainListRequest) (*types.QueryChainListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, err := k.GetAdminChainList(ctx, req.CreatorAddress)
	if !err {
		// data not found
		return &types.QueryChainListResponse{}, nil
	} else {
		return &types.QueryChainListResponse{
			ExelayerChains: val.Id,
		}, nil
	}

}
