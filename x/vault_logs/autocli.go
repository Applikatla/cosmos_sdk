package vaultlogs

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/x/vault_logs/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetAgentById",
					Use:       "get-agent [id]",
					Short:     "get agent data with id",
					Example:   "simd get-agent <agent_id>",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{
							ProtoField: "agent_id",
						},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "RegAgent",
					Use:       "reg-agent [agent_id] [agent_type]",
					Short:     "Register a new agent",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{
							ProtoField: "agent_id",
						},
						{
							ProtoField: "agent_type",
						},
					},
				},
			},
		},
	}
}
