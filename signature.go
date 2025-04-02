package nitrolite

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Signature represents an Ethereum signature in the R, S, V format.
// Compatible with Ethereum's EIP-155 transaction signatures.
type Signature struct {
	R [32]byte // First 32 bytes of the signature
	S [32]byte // Second 32 bytes of the signature
	V byte     // Recovery identifier (+27 per Ethereum convention)
}

// Sign hashes the provided data using Keccak256 with Ethereum's prefix and signs it with the given private key.
func Sign(data []byte, privateKey *ecdsa.PrivateKey) (Signature, error) {
	if privateKey == nil {
		return Signature{}, fmt.Errorf("private key is nil")
	}

	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	prefixedData := append([]byte(prefix), data...)

	dataHash := crypto.Keccak256Hash(prefixedData)
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
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	prefixedData := append([]byte(prefix), data...)
	dataHash := crypto.Keccak256Hash(prefixedData)

	signature := make([]byte, 65)
	copy(signature[0:32], sig.R[:])
	copy(signature[32:64], sig.S[:])

	if sig.V >= 27 {
		signature[64] = sig.V - 27
	}

	pubKeyRaw, err := crypto.Ecrecover(dataHash.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr == address, nil
}
