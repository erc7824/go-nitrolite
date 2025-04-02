package nitrolite

// State represents the current state of a channel.
type State struct {
	Data        []byte        // Arbitrary channel state information
	Allocations [2]Allocation // Fund distribution between participants
	Sigs        []Signature   // Signatures from channel participants
}
