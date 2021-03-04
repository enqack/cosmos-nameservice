package keeper

import (
	"github.com/enqack/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
