package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	vaultlogstypes "github.com/cosmos/cosmos-sdk/x/vault_logs/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	addressCodec address.Codec
	cdc          codec.BinaryCodec

	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema   collections.Schema
	Params   collections.Item[types.Params]
	RegAgent collections.Map[string, vaultlogstypes.AgentReg]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	addressCodec address.Codec,
	cdc codec.BinaryCodec,
	authority []byte,
) *Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}
	sb := collections.NewSchemaBuilder(storeService)

	k := &Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, types.ParamsKey, "vault_logs_param", codec.CollValue[types.Params](cdc)),
		RegAgent:     collections.NewMap(sb, vaultlogstypes.RegAgentKey, "vault_logs_agent_reg", collections.StringKey, codec.CollValue[vaultlogstypes.AgentReg](cdc)),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k
}

// logger returns module specific logs
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", "vault-logs"))
}

// Codec returns the module's codec (renamed for consistency)
func (k Keeper) Codec() codec.BinaryCodec {
	return k.cdc
}
