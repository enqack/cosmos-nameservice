package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/enqack/nameservice/x/nameservice/types"
)

// MinimumCreateWhoisPrice
func (k Keeper) CreateWhoisPrice(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyCreateWhoisPrice, &res)
	return
}

// UpdateWhoisPrice
func (k Keeper) UpdateWhoisPrice(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyUpdateWhoisPrice, &res)
	return
}

// DeleteWhoisPrice
func (k Keeper) DeleteWhoisPrice(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyDeleteWhoisPrice, &res)
	return
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.CreateWhoisPrice(ctx),
		k.UpdateWhoisPrice(ctx),
		k.DeleteWhoisPrice(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)

}
