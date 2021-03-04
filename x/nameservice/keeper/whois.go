package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/enqack/nameservice/x/nameservice/types"
	"strconv"
)

// GetWhoisCount get the total number of whois
func (k Keeper) GetWhoisCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisCountKey))
	byteKey := types.KeyPrefix(types.WhoisCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetWhoisCount set the total number of whois
func (k Keeper) SetWhoisCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisCountKey))
	byteKey := types.KeyPrefix(types.WhoisCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateWhois creates a whois with a new id and update the count
func (k Keeper) CreateWhois(ctx sdk.Context, msg types.MsgCreateWhois) {
	// Create the whois
	count := k.GetWhoisCount(ctx)
	var whois = types.Whois{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Name:    msg.Name,
		Address: msg.Address,
		Price:   msg.Price,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	key := types.KeyPrefix(types.WhoisKey + whois.Id)
	value := k.cdc.MustMarshalBinaryBare(&whois)
	store.Set(key, value)

	// Update whois count
	k.SetWhoisCount(ctx, count+1)
}

// SetWhois set a specific whois in the store
func (k Keeper) SetWhois(ctx sdk.Context, whois types.Whois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	b := k.cdc.MustMarshalBinaryBare(&whois)
	store.Set(types.KeyPrefix(types.WhoisKey+whois.Id), b)
}

// GetWhois returns a whois from its id
func (k Keeper) GetWhois(ctx sdk.Context, key string) types.Whois {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.WhoisKey+key)), &whois)
	return whois
}

// HasWhois checks if the whois exists
func (k Keeper) HasWhois(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	return store.Has(types.KeyPrefix(types.WhoisKey + id))
}

// GetWhoisOwner returns the creator of the whois
func (k Keeper) GetWhoisOwner(ctx sdk.Context, key string) string {
	return k.GetWhois(ctx, key).Creator
}

// DeleteWhois deletes a whois
func (k Keeper) DeleteWhois(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	store.Delete(types.KeyPrefix(types.WhoisKey + key))
}

// GetAllWhois returns all whois
func (k Keeper) GetAllWhois(ctx sdk.Context) (msgs []types.Whois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhoisKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Whois
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

//
// Functions used by querier
//

// Get creator of the item
func (k Keeper) GetCreator(ctx sdk.Context, key string) string {
	whois := k.GetWhois(ctx, key)
	return whois.Creator
}

// Check if the key exists in the store
func (k Keeper) Exists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisKey + key))
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	whois := k.GetWhois(ctx, name)
	return whois.Address
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, address string) {
	whois := k.GetWhois(ctx, name)
	whois.Address = address
	k.SetWhois(ctx, whois)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasCreator(ctx sdk.Context, name string) bool {
	whois := k.GetWhois(ctx, name)
	return len(whois.Creator) != 0
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetCreator(ctx sdk.Context, name string, creator string) {
	whois := k.GetWhois(ctx, name)
	whois.Creator = creator
	k.SetWhois(ctx, whois)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) string {
	whois := k.GetWhois(ctx, name)
	return whois.Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price string) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, whois)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.WhoisKey))
}

// Check if the key exists in the store
func (k Keeper) WhoisExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisKey + key))
}

