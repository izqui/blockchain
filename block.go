package main

type BlockSlice []Block
type Block struct {
	*BlockHeader
	*TransactionSlice
}

type BlockHeader struct {
	PrevBlock  []byte
	MerkelRoot []byte
	Timestamp  int32
	Nonce      int32
}

func NewBlock() Block {

	return Block{new(BlockHeader), new(TransactionSlice)}
}
