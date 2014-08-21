package main

import (
	"fmt"

	"github.com/izqui/helpers"
)

type TransactionsQueue chan *Transaction
type BlocksQueue chan *Block

type Blockchain struct {
	CurrentBlock Block
	BlockSlice
	TransactionsQueue
	BlocksQueue
}

func SetupBlockchan() *Blockchain {

	bl := new(Blockchain)
	bl.TransactionsQueue, bl.BlocksQueue = make(TransactionsQueue), make(BlocksQueue)

	bl.CurrentBlock = NewBlock()

	//Read blockchain from file and stuff...

	return bl
}

func (bl *Blockchain) Run() {

	for {
		select {
		case tr := <-bl.TransactionsQueue:
			if !self.CurrentBlock.TransactionSlice.Exists(*tr) {
				if tr.VerifyTransaction(helpers.ArrayOfBytes(TRANSACTION_POW_COMPLEXITY, POW_PREFIX)) {
					fmt.Println("Got a valid non existing transaction")
				}
			}
		}
	}
}
