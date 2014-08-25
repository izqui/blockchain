package main

import (
	"fmt"
	"time"

	"github.com/izqui/helpers"
)

type TransactionsQueue chan *Transaction
type BlocksQueue chan Block

type Blockchain struct {
	CurrentBlock Block
	BlockSlice
	TransactionsQueue
	BlocksQueue
}

func SetupBlockchan() *Blockchain {

	bl := new(Blockchain)
	bl.TransactionsQueue, bl.BlocksQueue = make(TransactionsQueue), make(BlocksQueue)

	//Read blockchain from file and stuff...

	bl.CreateNewBlock()

	return bl
}

func (bl *Blockchain) CreateNewBlock() {

	prevBlock := bl.BlockSlice.PreviousBlock()
	prevBlockHash := []byte{}
	if prevBlock != nil {

		prevBlockHash = prevBlock.Hash()
	}

	bl.CurrentBlock = NewBlock(prevBlockHash)
	bl.CurrentBlock.BlockHeader.Origin = self.Keypair.Public
}

func (bl *Blockchain) AddBlock(b Block) {

	bl.BlockSlice = append(bl.BlockSlice, b)
}

func (bl *Blockchain) Run() {

	interruptBlockGen := bl.GenerateBlocks()
	for {
		select {
		case tr := <-bl.TransactionsQueue:
			if !bl.CurrentBlock.TransactionSlice.Exists(*tr) {
				if tr.VerifyTransaction(TRANSACTION_POW) {

					bl.CurrentBlock.AddTransaction(tr)
					fmt.Println("new trans")

					interruptBlockGen <- true
				}
			}
		case b := <-bl.BlocksQueue:

			//verify block
			//Broadcast block and shit

			fmt.Println("New block!", b.TransactionSlice, b.VerifyBlock(BLOCK_POW))

			bl.AddBlock(b)
			bl.CreateNewBlock()

			interruptBlockGen <- true
		}
	}
}

func (bl *Blockchain) GenerateBlocks() chan bool {

	interrupt := make(chan bool)

	go func() {

	loop:
		block := bl.CurrentBlock
		block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
		block.BlockHeader.Nonce = 0
		block.BlockHeader.Timestamp = uint32(time.Now().Unix())

		for true {

			sleepTime := time.Microsecond
			if block.TransactionSlice.Len() > 0 {

				if CheckProofOfWork(BLOCK_POW, block.Hash()) {

					block.Signature = block.Sign(self.Keypair)
					bl.BlocksQueue <- block
					sleepTime = time.Hour * 24

				} else {

					block.BlockHeader.Nonce += 1
					fmt.Println(block.BlockHeader.Nonce)
				}

			} else {
				sleepTime = time.Hour * 24
			}

			select {
			case <-interrupt:
				goto loop
			case <-helpers.Timeout(sleepTime):
				continue
			}
		}
	}()

	return interrupt
}
