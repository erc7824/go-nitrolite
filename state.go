package nitrolite

// Signature represents an Ethereum signature
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// State represents the current state of a channel
type State struct {
	Data        []byte
	Allocations [2]Allocation
	Sigs        []Signature
}
