package main

import (
	"reflect"
	"testing"

	"github.com/izqui/helpers"
)

func TestKeyGeneration(t *testing.T) {

	keypair := GenerateNewKeypair(TEST_POW_PREFIX, TEST_KEY_POW_COMPLEXITY)

	if reflect.DeepEqual(keypair.Public[TEST_KEY_POW_COMPLEXITY:], helpers.ArrayOfBytes(TEST_KEY_POW_COMPLEXITY, TEST_POW_PREFIX)) {

		t.Error("Key creation error. It doesn't pass Proof of Work")
	}
}

func TestKeySigning(t *testing.T) {

	keypair := GenerateNewKeypair(TEST_POW_PREFIX, TEST_KEY_POW_COMPLEXITY)

	data := helpers.ArrayOfBytes(1024, 'a')
	signature, err := keypair.Sign(data)

	if err != nil {

		t.Error("base58 error")

	} else if !Verify(keypair.Public, signature, data) {

		t.Error("Signing and verifying error")
	}
}
