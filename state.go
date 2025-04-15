package nitrolite

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// State represents the current state of a channel.
type State struct {
	Data        []byte        // Arbitrary channel state information
	Allocations [2]Allocation // Fund distribution between participants
	Sigs        []Signature   // Signatures from channel participants
}

func EncodeState(channelID common.Hash, stateData []byte, allocations []Allocation) ([]byte, error) {
	allocationType, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{
			Name: "destination",
			Type: "address",
		},
		{
			Name: "token",
			Type: "address",
		},
		{
			Name: "amount",
			Type: "uint256",
		},
	})
	if err != nil {
		return nil, err
	}

	var allocValues []any
	for _, alloc := range allocations {
		allocValues = append(allocValues, struct {
			Destination common.Address
			Token       common.Address
			Amount      *big.Int
		}{
			Destination: alloc.Destination,
			Token:       alloc.Token,
			Amount:      alloc.Amount,
		})
	}

	args := abi.Arguments{
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // channelId
		{Type: abi.Type{T: abi.BytesTy}},                // data
		{Type: allocationType},                          // allocations as tuple[]
	}

	packed, err := args.Pack(channelID, stateData, allocations)
	if err != nil {
		return nil, err
	}
	return packed, nil
}
