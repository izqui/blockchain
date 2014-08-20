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
	tr.Signature = tr.Sign(kp)

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

func TestTransactionVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte("Hola que tal"))

	tr.Header.Nonce = tr.GenerateNonce(pow)
	tr.Signature = tr.Sign(kp)

	if !tr.VerifyTransaction(pow) {

		t.Error("Validation failing")
	}
}

func TestIncorrectPOWVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)
	powIncorrect := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, 'a')

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte("Hola que tal"))
	tr.Header.Nonce = tr.GenerateNonce(powIncorrect)
	tr.Signature = tr.Sign(kp)

	if tr.VerifyTransaction(pow) {

		t.Error("Passed validation without pow")
	}
}

func TestIncorrectSignatureVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)
	kp1, kp2 := GenerateNewKeypair(), GenerateNewKeypair()
	tr := NewTransaction(kp2.Public, nil, []byte("Hola que tal"))
	tr.Header.Nonce = tr.GenerateNonce(pow)
	tr.Signature = tr.Sign(kp1)

	if tr.VerifyTransaction(pow) {

		t.Error("Passed validation with incorrect key")
	}
}
