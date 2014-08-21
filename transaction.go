package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
	"time"

	"github.com/izqui/helpers"
)

type TransactionSlice []Transaction

func (slice TransactionSlice) Exists(tr Transaction) bool {

	for _, t := range slice {
		if reflect.DeepEqual(t.Signature, tr.Signature) {
			return true
		}
	}
	return false
}

func (slice TransactionSlice) AddTransaction(t Transaction) TransactionSlice {

	// Inserted sorted by timestamp
	for i, tr := range slice {
		if tr.Header.Timestamp >= t.Header.Timestamp {
			return append(append(slice[:i], t), slice[i:]...)
		}
	}

	return append(slice, t)
}

type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     uint32
	PayloadHash   []byte
	PayloadLength uint32
	Nonce         uint32
}

// Returns bytes to be sent to the network
func NewTransaction(from, to, payload []byte) *Transaction {

	t := Transaction{Header: TransactionHeader{From: from, To: to}, Payload: payload}

	t.Header.Timestamp = uint32(time.Now().Unix())
	t.Header.PayloadHash = helpers.SHA256(t.Payload)
	t.Header.PayloadLength = uint32(len(t.Payload))

	return &t
}
func (t *Transaction) Sign(keypair *Keypair) []byte {

	headerBytes, _ := t.Header.MarshalBinary()

	s, _ := keypair.Sign(helpers.SHA256(headerBytes))

	return s
}

func (t *Transaction) VerifyTransaction(pow []byte) bool {

	m, err := t.Header.MarshalBinary()
	if err != nil {
		return false
	}
	headerHash := helpers.SHA256(m)
	payloadHash := helpers.SHA256(t.Payload)

	return reflect.DeepEqual(payloadHash, t.Header.PayloadHash) && CheckProofOfWork(pow, headerHash) && SignatureVerify(t.Header.From, t.Signature, headerHash)
}

func (t *Transaction) GenerateNonce(prefix []byte) uint32 {

	newT := t
	for {
		header, _ := newT.Header.MarshalBinary()
		if CheckProofOfWork(prefix, helpers.SHA256(header)) {
			break
		}

		newT.Header.Nonce++
	}

	return newT.Header.Nonce
}

func CheckProofOfWork(prefix []byte, hash []byte) bool {

	if len(prefix) > 0 {
		return reflect.DeepEqual(prefix, hash[:len(prefix)])
	}
	return true
}

func (t *Transaction) MarshalBinary() ([]byte, error) {

	headerBytes, _ := t.Header.MarshalBinary()

	if len(headerBytes) != HEADER_SIZE {
		return nil, errors.New("Header marshalling error")
	}

	return append(append(headerBytes, helpers.FitBytesInto(t.Signature, NETWORK_KEY_SIZE)...), t.Payload...), nil
}

func (t *Transaction) UnmarshalBinary(d []byte) error {

	buf := bytes.NewBuffer(d)

	if len(d) < HEADER_SIZE+NETWORK_KEY_SIZE {
		return errors.New("Insuficient bytes for unmarshalling transaction")
	}

	header := &TransactionHeader{}
	if err := header.UnmarshalBinary(buf.Next(HEADER_SIZE)); err != nil {
		return err
	}

	t.Header = *header
	if len(d) != HEADER_SIZE+NETWORK_KEY_SIZE+int(t.Header.PayloadLength) {
		return errors.New("Payload length in header doesn't match with actual payload length")
	}

	t.Signature = helpers.StripByte(buf.Next(NETWORK_KEY_SIZE), 0)
	t.Payload = buf.Next(int(t.Header.PayloadLength))

	return nil

}

func (th *TransactionHeader) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	buf.Write(helpers.FitBytesInto(th.From, NETWORK_KEY_SIZE))
	buf.Write(helpers.FitBytesInto(th.To, NETWORK_KEY_SIZE))
	binary.Write(buf, binary.LittleEndian, th.Timestamp)
	buf.Write(helpers.FitBytesInto(th.PayloadHash, 32))
	binary.Write(buf, binary.LittleEndian, th.PayloadLength)
	binary.Write(buf, binary.LittleEndian, th.Nonce)

	return buf.Bytes(), nil

}

func (th *TransactionHeader) UnmarshalBinary(d []byte) error {

	buf := bytes.NewBuffer(d)
	th.From = helpers.StripByte(buf.Next(NETWORK_KEY_SIZE), 0)
	th.To = helpers.StripByte(buf.Next(NETWORK_KEY_SIZE), 0)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Timestamp)
	th.PayloadHash = buf.Next(32)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.PayloadLength)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Nonce)

	return nil
}
