package keeper

import (
	"context"

	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListExecutionLayers(goCtx context.Context, req *types.QueryListExecutionLayersRequest) (*types.QueryListExecutionLayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exelayers []types.Exelayer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	exelayerStore := prefix.NewStore(store, types.KeyPrefix(types.ExelayerKeyPrefix))

	pageRes, err := query.Paginate(exelayerStore, req.Pagination, func(key []byte, value []byte) error {
		_ = key
		var exelayer types.Exelayer
		if err := k.cdc.Unmarshal(value, &exelayer); err != nil {
			return err
		}
		exelayers = append(exelayers, exelayer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListExecutionLayersResponse{Exelayer: exelayers, Pagination: pageRes}, nil
}
