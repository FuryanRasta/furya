// nolint

package poe

import (
	"github.com/furyanrasta/furya/x/poe/keeper"
	"github.com/furyanrasta/furya/x/poe/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

type DeliverTxfn = keeper.DeliverTxFn
