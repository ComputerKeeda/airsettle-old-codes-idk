package keeper

import (
	"airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetExecutionlayers(ctx sdk.Context, exelayer types.Exelayer) {

	// store verification key seperately. save memory on processing.
	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	vkey_formate := types.Vkey{
		Id:              exelayer.Id,
		VerificationKey: exelayer.VerificationKey,
	}
	vk_binary := k.cdc.MustMarshal(&vkey_formate)
	vk_store.Set([]byte(exelayer.Id), vk_binary)

	// store other details. except vKey
	exelayer.VerificationKey = "check this out" //"/getvkey/" + exelayer.Id + "/"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerKeyPrefix))
	b := k.cdc.MustMarshal(&exelayer)
	store.Set([]byte(exelayer.Id), b)

	// //* update/create chainlist. e.g. chains under admin/creator/address
	val, found := k.GetAdminChainList(ctx, exelayer.Creator)
	if !found {
		// key don't exists, create new list
		var chain_list = types.ExelayerChains{
			Creator: exelayer.Creator,
			Id:      []string{exelayer.Id},
		}
		// create admin chain list
		k.SetAdminChainList(ctx, chain_list)
	} else {
		// key exists, append value & save value
		val.Id = append(val.Id, exelayer.Id)
		k.SetAdminChainList(ctx, val)
	}
}

func (k Keeper) UpdateExecutionlayers(ctx sdk.Context, exelayer types.Exelayer) {

	// store other details. except vKey
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerKeyPrefix))
	b := k.cdc.MustMarshal(&exelayer)
	store.Set([]byte(exelayer.Id), b)

	// //* update/create chainlist. e.g. chains under admin/creator/address
	val, found := k.GetAdminChainList(ctx, exelayer.Creator)
	if !found {
		// key don't exists, create new list
		var chain_list = types.ExelayerChains{
			Creator: exelayer.Creator,
			Id:      []string{exelayer.Id},
		}
		// create admin chain list
		k.SetAdminChainList(ctx, chain_list)
	} else {
		// key exists, append value & save value
		val.Id = append(val.Id, exelayer.Id)
		k.SetAdminChainList(ctx, val)
	}
}

func (k Keeper) SetAdminChainList(ctx sdk.Context, chain_list types.ExelayerChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerChainKey))
	appendedValue := k.cdc.MustMarshal(&chain_list)
	store.Set([]byte(chain_list.Creator), appendedValue)
}

func (k Keeper) GetAdminChainList(
	ctx sdk.Context,
	creator string,
) (val types.ExelayerChains, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerChainKey))
	b := store.Get([]byte(creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetExecutionlayers(ctx sdk.Context, id string) (val types.Exelayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerKeyPrefix))

	b := store.Get(types.ExelayerKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetVerificationKey(ctx sdk.Context, id string) (val types.Vkey, found bool) {
	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	b := vk_store.Get(types.ExelayerKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetExelayerById(ctx sdk.Context, id string) (val types.Exelayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExelayerKeyPrefix))

	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// func (k Keeper) LocalGetVkey(ctx sdk.Context, chainid string) string {
// 	// get verification key
// 	var vk_val types.Vkey
// 	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
// 	b := vk_store.Get([]byte(chainid))
// 	if b == nil {
// 		vk_val.VerificationKey = ""
// 		// verification key not found
// 	} else {
// 		k.cdc.MustUnmarshal(b, &vk_val)
// 	}
// 	return vk_val.VerificationKey
// }
