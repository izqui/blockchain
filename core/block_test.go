package core

import (
	"reflect"
	"testing"

	"github.com/izqui/helpers"
)

func TestMerkellHash(t *testing.T) {

	tr1 := NewTransaction(nil, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024*1024))))
	tr2 := NewTransaction(nil, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024*1024))))
	tr3 := NewTransaction(nil, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024*1024))))
	tr4 := NewTransaction(nil, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024*1024))))

	b := new(Block)
	b.TransactionSlice = &TransactionSlice{*tr1, *tr2, *tr3, *tr4}

	mt := b.GenerateMerkelRoot()
	manual := helpers.SHA256(append(helpers.SHA256(append(tr1.Hash(), tr2.Hash()...)), helpers.SHA256(append(tr3.Hash(), tr4.Hash()...))...))

	if !reflect.DeepEqual(mt, manual) {
		t.Error("Merkel tree generation fails")
	}
}

//TODO: Write block validation and marshalling tests [Issue: https://github.com/izqui/blockchain/issues/2]

/*
func TestBlockMarshalling(t *testing.T) {

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024*1024))))

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

func TestBlockVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024))))

	tr.Header.Nonce = tr.GenerateNonce(pow)
	tr.Signature = tr.Sign(kp)

	if !tr.VerifyTransaction(pow) {

		t.Error("Validation failing")
	}
}

func TestIncorrectBlockPOWVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)
	powIncorrect := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, 'a')

	kp := GenerateNewKeypair()
	tr := NewTransaction(kp.Public, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024))))
	tr.Header.Nonce = tr.GenerateNonce(powIncorrect)
	tr.Signature = tr.Sign(kp)

	if tr.VerifyTransaction(pow) {

		t.Error("Passed validation without pow")
	}
}

func TestIncorrectBlockSignatureVerification(t *testing.T) {

	pow := helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, TEST_POW_PREFIX)
	kp1, kp2 := GenerateNewKeypair(), GenerateNewKeypair()
	tr := NewTransaction(kp2.Public, nil, []byte(helpers.RandomString(helpers.RandomInt(0, 1024))))
	tr.Header.Nonce = tr.GenerateNonce(pow)
	tr.Signature = tr.Sign(kp1)

	if tr.VerifyTransaction(pow) {

		t.Error("Passed validation with incorrect key")
	}
}
*/
