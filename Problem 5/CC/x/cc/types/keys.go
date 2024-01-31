package types

const (
	// ModuleName defines the module name
	ModuleName = "cc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cc"
)

var (
	ParamsKey = []byte("p_cc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
