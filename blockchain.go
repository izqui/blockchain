package main

import (
	"fmt"
	"reflect"
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

			if bl.CurrentBlock.TransactionSlice.Exists(*tr) {
				continue
			}
			if !tr.VerifyTransaction(TRANSACTION_POW) {
				fmt.Println("Recieved non valid transaction", tr)
				continue
			}

			bl.CurrentBlock.AddTransaction(tr)
			interruptBlockGen <- bl.CurrentBlock

			//Broadcast transaction to the network
			mes := NewMessage(MESSAGE_SEND_TRANSACTION)
			mes.Data, _ = tr.MarshalBinary()

			self.Network.BroadcastQueue <- *mes

		case b := <-bl.BlocksQueue:

			if !bl.BlockSlice.Exists(b) {
				if b.VerifyBlock(BLOCK_POW) {

					if reflect.DeepEqual(b.PrevBlock, bl.CurrentBlock.Hash()) {
						// I'm missing some blocks in the middle. Request'em.
						fmt.Println("Missing blocks in between")
					} else {

						fmt.Println("New block!", b.TransactionSlice, b.VerifyBlock(BLOCK_POW))

						transDiff := TransactionSlice{}

						if !reflect.DeepEqual(b.BlockHeader.MerkelRoot, bl.CurrentBlock.MerkelRoot) {
							// Transactions are different
							fmt.Println("Transactions are different. finding diff")
							transDiff = DiffTransactionSlices(*bl.CurrentBlock.TransactionSlice, *b.TransactionSlice)
						}

						bl.AddBlock(b)
						bl.CreateNewBlock()
						bl.CurrentBlock.TransactionSlice = &transDiff

						interruptBlockGen <- bl.CurrentBlock

						//Broadcast block and shit
					}
				}
			}
		}
	}
}

func DiffTransactionSlices(a, b TransactionSlice) (diff TransactionSlice) {
	//Assumes transaction arrays are sorted (which maybe is too big of an assumption)
	lastj := 0
	for _, t := range a {
		found := false
		for j := lastj; j < len(b); j++ {
			if reflect.DeepEqual(b[j].Signature, t.Signature) {
				found = true
				lastj = j
				break
			}
		}
		if !found {
			diff = append(diff, t)
		}
	}

	return
}

func (bl *Blockchain) GenerateBlocks() chan Block {

	interrupt := make(chan Block)

	go func() {

		block := <-interrupt
	loop:
		fmt.Println("Starting Proof of Work...")
		block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
		block.BlockHeader.Nonce = 0
		block.BlockHeader.Timestamp = uint32(time.Now().Unix())

		for true {

			sleepTime := time.Nanosecond
			if block.TransactionSlice.Len() > 0 {

				if CheckProofOfWork(BLOCK_POW, block.Hash()) {

					block.Signature = block.Sign(self.Keypair)
					bl.BlocksQueue <- block
					sleepTime = time.Hour * 24

				} else {

					block.BlockHeader.Nonce += 1
				}

			} else {
				sleepTime = time.Hour * 24
			}

			select {
			case block = <-interrupt:
				goto loop
			case <-helpers.Timeout(sleepTime):
				continue
			}
		}
	}()

	return interrupt
}
