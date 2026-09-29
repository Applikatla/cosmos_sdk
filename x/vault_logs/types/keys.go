package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "vault_logs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vault_logs"
)

var (
	RegAgentKey = collections.NewPrefix(ModuleName + "_reg_agent")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
