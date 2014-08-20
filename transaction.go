package main

import (
	"errors"
	"github.com/izqui/helpers"
	"time"
)

type TransactionSlice []Transaction

type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     int
	PayloadHash   []byte
	PayloadLength int
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

	if len(headerBytes) != HEADER_SIZE {
		return nil, errors.New("Header marshalling error")
	}

	return append(append(headerBytes, helpers.FitBytesInto(t.Signature, NETWORK_KEY_SIZE)...), t.Payload...), nil
}

func (t *Transaction) UnmarshalBinary([]byte) {

}

func (th TransactionHeader) MarshalBinary() ([]byte, error) {

	return []byte{}, nil
}

func (th TransactionHeader) UnmarshalBinary([]byte) {

}
