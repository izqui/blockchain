package main

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

	for i := 0; i < 50; i++ {
		keypair := GenerateNewKeypair()

		data := helpers.ArrayOfBytes(i, 'a')
		signature, err := keypair.Sign(data)

		if err != nil {

			t.Error("base58 error")

		} else if !SignatureVerify(keypair.Public, signature, data) {

			t.Error("Signing and verifying error", len(keypair.Public))
		}
	}

}
