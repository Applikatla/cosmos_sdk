package vaultlogs

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/vault_logs/keeper"
	"github.com/cosmos/cosmos-sdk/x/vault_logs/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const ConsensusVersion = 5

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	// _ module.AppModuleSimulation = AppModule{}
	// _ module.HasGenesis          = AppModule{}
	_ module.HasServices = AppModule{}

	_ appmodule.AppModule     = AppModule{}
	_ appmodule.HasEndBlocker = AppModule{}
)

// AppModuleBasic defines the basic application module used by the gov module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
	ac  address.Codec
}

// app module basic constructor
func NewAppModuleBasic(cdc codec.BinaryCodec) *AppModuleBasic {
	return &AppModuleBasic{
		cdc: cdc,
	}
}

// Name returns the bank module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// return cdc.MustMarshalJSON(types.DefaultGenesis())
	return nil
}

// func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
// 	var genState types.GenesisState
// 	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
// 		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
// 	}
// 	return genState.Validate()
// }

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// app module

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: *NewAppModuleBasic(cdc),
		keeper:         keeper,
	}
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) EndBlock(ctx context.Context) error {
	sdkCtx := ctx.(sdk.Context)
	_ = sdkCtx
	// EndBlocker(sdkCtx, am.keeper)
	return nil
}

func (am AppModule) IsAppModule() {

}

func (am AppModule) IsOnePerModuleType() {

}
