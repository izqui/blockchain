package main

import (
	"bytes"

	"github.com/izqui/functional"
	"github.com/izqui/helpers"
)

type BlockSlice []Block
type Block struct {
	*BlockHeader
	*TransactionSlice
}

type BlockHeader struct {
	Origin     []byte
	PrevBlock  []byte
	MerkelRoot []byte
	Timestamp  int32
	Nonce      int32
}

func NewBlock() Block {

	return Block{new(BlockHeader), new(TransactionSlice)}
}

func (b *Block) GenerateMerkelRoot() []byte {

	var merkell func(hashes [][]byte) []byte
	merkell = func(hashes [][]byte) []byte {

		l := len(hashes)
		if l == 1 {
			return hashes[0]
		} else {

			if l%2 == 1 {
				return merkell([][]byte{merkell(hashes[:l-1]), hashes[l-1]})
			}

			bs := make([][]byte, l/2)
			for i, _ := range bs {
				i, j := i*2, (i*2)+1
				bs[i] = helpers.SHA256(append(hashes[i], hashes[j]...))
			}
			return merkell(bs)
		}
	}

	return merkell(functional.Map(func(t Transaction) []byte { return t.Hash() }, []Transaction(*b.TransactionSlice)).([][]byte))

	return nil
}
func (b *Block) MarshalBinary() ([]byte, error) {

	bhb, err := b.BlockHeader.MarshalBinary()
	if err != nil {
		return nil, err
	}

	tsb, err := b.TransactionSlice.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return append(bhb, tsb...), nil
}

func (b *Block) UnmarshalBinary(d []byte) error {

	/*buf := bytes.NewBuffer(d)

	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Timestamp)
	th.PayloadHash = buf.Next(32)*/

	return nil
}

func (h *BlockHeader) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)
	/*
		binary.Write(buf, binary.LittleEndian, th.Timestamp)
		buf.Write(helpers.FitBytesInto(th.PayloadHash, 32))
	*/
	return buf.Bytes(), nil
}

func (h *BlockHeader) UnmarshalBinary(d []byte) error {

	/*buf := bytes.NewBuffer(d)

	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Timestamp)
	th.PayloadHash = buf.Next(32)*/

	return nil
}
