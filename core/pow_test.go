package core

import (
	"testing"
)

func TestPow(t *testing.T) {

	b1 := CheckProofOfWork([]byte{0, 0, 0, 1, 2, 3}, []byte{0, 0, 0, 1, 2, 3, 4, 5})
	b2 := CheckProofOfWork([]byte{0, 0}, []byte("hola"))
	b3 := CheckProofOfWork(BLOCK_POW, append(BLOCK_POW, 1))
	b4 := CheckProofOfWork(nil, []byte("hola que tal"))

	if !b1 || b2 || !b3 || !b4 {
		t.Error("Proof of work test fails.")
	}
}
