package main

import (
	"fmt"
	"github.com/izqui/helpers"
	"reflect"
	"testing"
)

func TestKeyGeneration(t *testing.T) {

	keypair := GenerateNewKeypair(TEST_POW_PREFIX, TEST_KEY_POW_COMPLEXITY)

	if reflect.DeepEqual(keypair.Public[TEST_KEY_POW_COMPLEXITY:], helpers.ArrayOfBytes(TEST_KEY_POW_COMPLEXITY, TEST_POW_PREFIX)) {

		t.Error("Key creation error. It doesn't pass Proof of Work")
	}
}

func TestKeySigning(t *testing.T) {

	for i := 0; i < 5000; i++ {
		keypair := GenerateNewKeypair(TEST_POW_PREFIX, TEST_KEY_POW_COMPLEXITY)

		data := helpers.ArrayOfBytes(i, i)
		signature, err := keypair.Sign(data)

		if err != nil {

			t.Error("base58 error")

		} else if !Verify(keypair.Public, signature, data) {

			t.Error("Signing and verifying error", len(keypair.Public))
		}
	}

}
