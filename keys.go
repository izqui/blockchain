package main

import (
	_ "fmt"
	"reflect"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	_ "math/big"

	"github.com/izqui/helpers"
	"github.com/tv42/base58"
)

// Key generation with proof of work
type Keypair struct {
	Public  []byte `json:"public"`  //x (base58 encoded) + y (base58 encoded)
	Private []byte `json:"private"` //d (base58 encoded)
}

func randomKeyPair() Keypair {

	pk, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)

	public := base58.EncodeBig(base58.EncodeBig([]byte{}, pk.PublicKey.X), pk.PublicKey.Y)
	private := base58.EncodeBig([]byte{}, pk.D)

	kp := Keypair{Public: public, Private: private}

	return kp
}

func GenerateNewKeypair(prefix byte, complexity int) *Keypair {

	pr := helpers.ArrayOfBytes(complexity, prefix)

	var kp Keypair

	for {

		kp = randomKeyPair()

		//Proof of work
		if complexity == 0 || reflect.DeepEqual(kp.Public[:complexity], pr) {
			break
		}
	}
	return &kp
}

func (k *Keypair) Sign(data []byte) ([]byte, error) {

	hash := helpers.SHA256(data)
	d, err := base58.DecodeToBig(k.Private)
	if err != nil {
		return nil, err
	}

	publicKey := k.Public
	keyLength := len(publicKey) / 2

	x, err := base58.DecodeToBig(publicKey[:keyLength])
	if err != nil {
		return nil, err
	}
	y, err := base58.DecodeToBig(publicKey[keyLength:])
	if err != nil {
		return nil, err
	}

	key := ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P224(), x, y}, d}

	r, s, _ := ecdsa.Sign(rand.Reader, &key, hash)

	return base58.EncodeBig(base58.EncodeBig([]byte{}, r), s), nil
}

func Verify(publicKey []byte, sig, data []byte) bool {

	hash := helpers.SHA256(data)

	keyLength := len(publicKey) / 2
	x, err := base58.DecodeToBig(publicKey[:keyLength])
	if err != nil {

		return false
	}
	y, err := base58.DecodeToBig(publicKey[keyLength:])
	if err != nil {

		return false
	}

	sigLength := len(sig) / 2

	r, err := base58.DecodeToBig(sig[:sigLength])
	if err != nil {

		return false
	}
	s, err := base58.DecodeToBig(sig[sigLength:])
	if err != nil {

		return false
	}

	pub := ecdsa.PublicKey{elliptic.P224(), x, y}

	return ecdsa.Verify(&pub, hash, r, s)
}
