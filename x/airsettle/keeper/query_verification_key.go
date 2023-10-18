package keeper

import (
	"context"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerificationKey(goCtx context.Context, req *types.QueryVerificationKeyRequest) (*types.QueryVerificationKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	vkey := k.LocalGetvkey(ctx, req.Id)
	return &types.QueryVerificationKeyResponse{
		Vkey: vkey,
	}, nil
}
