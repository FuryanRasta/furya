package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	poewasm "github.com/furyanrasta/furya/x/poe/wasm"
	twasmkeeper "github.com/furyanrasta/furya/x/twasm/keeper"
	twasmtypes "github.com/furyanrasta/furya/x/twasm/types"
)

func SetupWasmHandlers(
	cdc codec.Codec,
	bankKeeper twasmtypes.BankKeeper,
	govRouter govtypes.Router,
	twasmKeeper twasmkeeper.FuryaWasmHandlerKeeper,
	poeKeeper poewasm.ViewKeeper,
	consensusParamsUpdater twasmkeeper.ConsensusParamsUpdater,
) []wasmkeeper.Option {
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Staking: poewasm.StakingQuerier(poeKeeper),
		Custom:  poewasm.CustomQuerier(poeKeeper),
	})

	extMessageHandlerOpt := wasmkeeper.WithMessageHandlerDecorator(func(nested wasmkeeper.Messenger) wasmkeeper.Messenger {
		return wasmkeeper.NewMessageHandlerChain(
			// disable staking messages
			wasmkeeper.MessageHandlerFunc(func(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) (events []sdk.Event, data [][]byte, err error) {
				if msg.Staking != nil {
					return nil, nil, sdkerrors.Wrap(wasmtypes.ErrExecuteFailed, "not supported, yet")
				}
				return nil, nil, wasmtypes.ErrUnknownMsg
			}),
			nested,
			// append our custom message handler
			twasmkeeper.NewFuryaHandler(cdc, twasmKeeper, bankKeeper, consensusParamsUpdater, govRouter),
		)
	})
	return []wasm.Option{
		queryPluginOpt,
		extMessageHandlerOpt,
	}
}
