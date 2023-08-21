package keeper

import (
	"context"

	"github.com/airchains-network/Airchains-settlement-layer/x/airsettle/types"
	// verifier "airsettle/x/airsettle/verifier"

	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Verify(goCtx context.Context, req *types.QueryVerifyRequest) (*types.QueryVerifyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get verification key
	str_verificationKey := k.LocalGetvkey(ctx, req.Id)
	if str_verificationKey == "" {
		Log("verification key not found")
		return &types.QueryVerifyResponse{
			Result:  false,
			Message: "verification key not found",
		}, sdkerrors.ErrKeyNotFound
	}

	batch, found := k.GetbatchById(ctx, req.Id, req.BatchNumber)
	str_zkproof := batch.ZkProof
	if !found {
		Log("Proof not found")
		return &types.QueryVerifyResponse{
			Result:  false,
			Message: "proof not found",
		}, sdkerrors.ErrKeyNotFound
	}

	Log("verify request received")

	// verify
	result, message := k.Verifier(ctx, str_zkproof, str_verificationKey, req.Inputs)

	return &types.QueryVerifyResponse{
		Result:  result,
		Message: message,
	}, nil
}
