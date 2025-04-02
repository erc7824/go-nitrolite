package nitrolite

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Signature represents an Ethereum signature in the standard V, R, S format.
// Compatible with Ethereum's EIP-155 transaction signatures.
type Signature struct {
	V uint8    // Recovery identifier (+27 per Ethereum convention)
	R [32]byte // First 32 bytes of the signature
	S [32]byte // Second 32 bytes of the signature
}

// Sign hashes the provided data using Keccak256 and signs it with the given private key.
// Returns an Ethereum-compatible signature that can be verified by Solidity contracts.
func Sign(data []byte, privateKey *ecdsa.PrivateKey) (Signature, error) {
	dataHash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		return Signature{}, fmt.Errorf("failed to sign data: %w", err)
	}

	if len(signature) != 65 {
		return Signature{}, fmt.Errorf("invalid signature length: got %d, want 65", len(signature))
	}

	var sig Signature
	copy(sig.R[:], signature[:32])
	copy(sig.S[:], signature[32:64])
	sig.V = signature[64] + 27

	return sig, nil
}

// Verify checks if the signature on the provided data was created by the given address.
func Verify(data []byte, sig Signature, address common.Address) (bool, error) {
	dataHash := crypto.Keccak256Hash(data)

	signature := make([]byte, 65)
	copy(signature[0:32], sig.R[:])
	copy(signature[32:64], sig.S[:])
	signature[64] = sig.V - 27

	pubKeyRaw, err := crypto.Ecrecover(dataHash.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	pubKeyBytes := pubKeyRaw[1:] // Remove the compression indicator byte
	pubKeyHash := crypto.Keccak256(pubKeyBytes)
	recoveredAddr := common.BytesToAddress(pubKeyHash[12:]) // Last 20 bytes

	return recoveredAddr == address, nil
}
