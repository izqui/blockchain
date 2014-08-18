package main

type Block struct {
	BlockHeader
	BlockBody
}

type BlockHeader struct {
	PrevBlock  []byte
	MerkelRoot []byte
	Timestamp  int
}

type BlockBody struct {
	TransactionSlice
}
