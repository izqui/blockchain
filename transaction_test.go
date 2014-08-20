package main

import (
	"github.com/izqui/helpers"
	"reflect"
	"testing"
)

func TestTransactionMarshalling(t *testing.T) {

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte("Hola que tal"))

	tr.Header.Nonce = tr.GenerateNonce(helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX))
	tr.Signature = tr.Sign(*kp)

	data, err := tr.MarshalBinary()

	if err != nil {
		t.Error(err)
	}

	newT := &Transaction{}
	err = newT.UnmarshalBinary(data)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(*newT, *tr) {
		t.Error("Marshall, unmarshall failed")
	}

}
