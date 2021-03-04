package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/enqack/nameservice/x/nameservice/keeper"
	"github.com/enqack/nameservice/x/nameservice/types"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateWhois) (*sdk.Result, error) {
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

	// Checks that the element exists
	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWhois(ctx, whois)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteWhois(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteWhois) (*sdk.Result, error) {
	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteWhois(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
