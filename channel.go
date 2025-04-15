package nitrolite

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

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
