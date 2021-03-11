package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/enqack/nameservice/x/nameservice/keeper"
	"github.com/enqack/nameservice/x/nameservice/types"
)

var (
	minimumCreateWhoisPrice, _ = types.CoinsFromString("10trycoin")
	updateWhoisPrice, _        = types.CoinsFromString("5trycoin")
	deleteWhoisPrice, _        = types.CoinsFromString("1trycoin")
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateWhois) (*sdk.Result, error) {
	// Check if whois name already exists
	if k.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name already exists")
	}

	// Check is name is valid
	if !k.VerifyNameFormat(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name is not valid")
	}

	// Check if address is valid
	if !k.IsValidAddress(ctx, msg.Address) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address is not valid")
	}

	// Convert creator (type string) to sdk.AccAddress type
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Convert price (type string) to sdk.Coins type
	price, err := types.CoinsFromString(msg.Price)
	if err != nil {
		return nil, err
	}

	// Check if price is above minimum
	if minimumCreateWhoisPrice.IsAllGT(price) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "price below minimum")
	}

	// Deduct coins from creator's account
	err = k.CoinKeeper.SubtractCoins(ctx, creator, price)
	if err != nil {
		return nil, err
	}

	k.CreateWhois(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateWhois) (*sdk.Result, error) {
	var whois = types.Whois{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Address: msg.Address,
		Price:   msg.Price,
	}

	// Check that the element exists
	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Check if the the msg sender is the same as the current owner
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if whois name already exists
	if k.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name already exists")
	}

	// Check if name is valid
	if !k.VerifyNameFormat(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name is not valid")
	}

	// Check if address is valid
	if !k.IsValidAddress(ctx, msg.Address) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address is not valid")
	}

	// Convert creator (type string) to sdk.AccAddress type
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Dedeuct coins from owner's account
	err = k.CoinKeeper.SubtractCoins(ctx, creator, updateWhoisPrice)
	if err != nil {
		return nil, err
	}

	k.SetWhois(ctx, whois)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteWhois) (*sdk.Result, error) {
	// Check if id exists
	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Check if the the msg sender is the same as the current owner
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Convert creator (type string) to sdk.AccAddress type
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Dedeuct coins from owner's account
	err = k.CoinKeeper.SubtractCoins(ctx, creator, deleteWhoisPrice)
	if err != nil {
		return nil, err
	}

	k.DeleteWhois(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
