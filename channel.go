package nitrolite

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

// GetChannelID returns the keccak256 hash of the ABI-encoded channel data.
// The encoding packs the two participants, the adjudicator, the challenge, and the nonce
// as static types (addresses padded to 32 bytes, and uint64 values in a 32-byte big-endian form).
func GetChannelID(ch Channel) common.Hash {
	var data []byte

	// Encode the two participants: pad each address to 32 bytes.
	for _, addr := range ch.Participants {
		paddedAddr := common.LeftPadBytes(addr.Bytes(), 32)
		data = append(data, paddedAddr...)
	}

	// Encode the adjudicator address.
	data = append(data, common.LeftPadBytes(ch.Adjudicator.Bytes(), 32)...)

	// Encode the challenge (uint64) into 32 bytes (big-endian).
	challengeBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(challengeBytes[24:], ch.Challenge)
	data = append(data, challengeBytes...)

	// Encode the nonce (uint64) into 32 bytes (big-endian).
	nonceBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(nonceBytes[24:], ch.Nonce)
	data = append(data, nonceBytes...)

	// Compute and return the keccak256 hash of the concatenated data.
	return crypto.Keccak256Hash(data)
}
