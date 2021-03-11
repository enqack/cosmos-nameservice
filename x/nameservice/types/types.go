package types

import (
	"regexp"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


// CoinsFromString - convert price (type string) in the form of <amount><denom> to type sdk.Coins
func CoinsFromString(price string) (sdk.Coins, error) {
	re := regexp.MustCompile("[[:digit:]]+")
	// get location of price amount
	loc := re.FindIndex([]byte(price))
	// get price amount as string
	strAmount := re.FindString(price)
	// get price demon from location of amount to end of price string 
	denom := price[loc[1]:len(price)]
	// convert to lower case
	denom = strings.ToLower(denom)
	// convert price amount from string to int64
	amount, err := strconv.ParseInt(strAmount, 10, 64)
	if err != nil {
		return nil, err
	}
	// return Coins type
	return sdk.NewCoins(sdk.NewInt64Coin(denom, amount)), nil
}
