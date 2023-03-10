package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = wasmtypes.ModuleName

	// RouterKey is used to route governance proposals
	RouterKey = wasmtypes.RouterKey

	// StoreKey is the prefix under which we store this module's data
	StoreKey = wasmtypes.StoreKey
)

// nolint
var (
	// leave enough space to not conflict with wasm prefixes

	PrivilegedContractsSecondaryIndexPrefix = []byte{0xa0}
	ContractCallbacksSecondaryIndexPrefix   = []byte{0xa1}
)
