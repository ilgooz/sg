package keeper

import (
	"github.com/test/sg/x/sg/types"
)

var _ types.QueryServer = Keeper{}
