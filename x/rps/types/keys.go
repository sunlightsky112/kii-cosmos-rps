package types

// Here store the stateless collection keys
var ()

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ModuleName = "rps"
	RouterKey  = ModuleName // used for message routing
	StoreKey   = ModuleName // used for KVStore prefix
	StudentKey = "Student/value/"
)
