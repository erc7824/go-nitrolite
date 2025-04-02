package nitrolite

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestSignAndVerify(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	testData := []byte("test signature data")

	// Test valid signature
	sig, err := Sign(testData, privateKey)
	if err != nil {
		t.Fatalf("Failed to sign data: %v", err)
	}
	valid, err := Verify(testData, sig, address)
	if err != nil || !valid {
		t.Fatal("Signature verification failed with correct address")
	}

	// Test wrong address
	wrongKey, _ := crypto.GenerateKey()
	wrongAddress := crypto.PubkeyToAddress(wrongKey.PublicKey)
	valid, err = Verify(testData, sig, wrongAddress)
	if err != nil || valid {
		t.Fatal("Signature verification should have failed with wrong address")
	}

	// Test modified data
	modifiedData := []byte("modified test data")
	valid, err = Verify(modifiedData, sig, address)
	if err != nil || valid {
		t.Fatal("Signature verification should have failed with modified data")
	}
}
