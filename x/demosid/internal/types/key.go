package types

const (
	// ModuleName is the name of the module
	ModuleName = "demosid"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// Constant for the msgs header
	AttributeBytesHeader              = []byte{0xD1}
	DataAccessGrantRequestBytesHeader = []byte{0xD2}

	DefaultKeySeparator = []byte{0xFF, 0xDD, 0x00}
)
