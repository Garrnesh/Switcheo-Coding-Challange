package types

const (
	// ModuleName defines the module name
	ModuleName = "prob"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_prob"
)

var (
	ParamsKey = []byte("p_prob")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
