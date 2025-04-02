package nitrolite

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Allocation represents a fund allocation in a channel
type Allocation struct {
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
}

// Channel represents the configuration of a state channel
type Channel struct {
	Participants [2]common.Address
	Adjudicator  common.Address
	Challenge    uint64
	Nonce        uint64
}

