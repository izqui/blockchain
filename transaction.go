package main

import (
	"time"
)

type TransactionSlice []Transaction

type Transaction struct {
	Header    TransactionHeader
	Payload   []byte
	Signature []byte
}

type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     int
	ContentLength int
	ContentHash   []byte
}

// Returns bytes to be sent to the network

func (t Transaction) Sign(keypair Keypair) Transaction {

	newT := t
	newT.Header.Timestamp = int(time.Now().Unix())

	headerBytes, _ := t.Header.MarshalBinary()
	newT.Signature, _ = keypair.Sign(headerBytes)
	return newT
}

func (t *Transaction) MarshalBinary() ([]byte, error) {

	headerBytes, _ := t.Header.MarshalBinary()
	return (append(append(headerBytes, t.Payload...), t.Signature...)), nil
}

func (t *Transaction) UnmarshalBinary() []byte {
	return []byte{}
}

func (th TransactionHeader) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

func (th TransactionHeader) UnmarshalBinary() []byte {
	return []byte{}
}
