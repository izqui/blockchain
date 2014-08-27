package core

import (
	"github.com/izqui/helpers"

	"testing"
)

func TestKeyGeneration(t *testing.T) {

	keypair := GenerateNewKeypair()

	if len(keypair.Public) > 80 {
		t.Error("Error generating key")
	}
}

func TestKeySigning(t *testing.T) {

	for i := 0; i < 5; i++ {
		keypair := GenerateNewKeypair()

		data := helpers.ArrayOfBytes(i, 'a')
		hash := helpers.SHA256(data)

		signature, err := keypair.Sign(hash)

		if err != nil {

			t.Error("base58 error")

		} else if !SignatureVerify(keypair.Public, signature, hash) {

			t.Error("Signing and verifying error", len(keypair.Public))
		}
	}

}
