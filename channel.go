package nitrolite

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Allocation represents a fund allocation in a state channel.
type Allocation struct {
	Destination common.Address // Recipient of the funds
	Token       common.Address // ERC-20 token address (zero address for ETH)
	Amount      *big.Int       // Quantity of tokens allocated
}

// Channel represents the configuration of a state channel between two participants.
type Channel struct {
	Participants [2]common.Address // Addresses of the two channel parties
	Adjudicator  common.Address    // Address of the dispute resolution contract
	Challenge    uint64            // Challenge period duration in seconds
	Nonce        uint64            // Unique identifier for this channel
}
