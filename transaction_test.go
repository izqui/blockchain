package main

import (
	"fmt"
	"github.com/izqui/helpers"
	"testing"
)

func TestTransactionMarshalling(t *testing.T) {

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte("Hola que tal"))

	tr.Header.Nonce = tr.ProofOfWork(helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX))
	tr.Signature = tr.Sign(*kp)

	_, err := tr.MarshalBinary()

	if err != nil {
		t.Error(err)
	}
}
