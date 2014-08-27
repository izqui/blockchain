package core

import (
	"fmt"
	"log"
)

var Core = struct {
	*Keypair
	*Blockchain
	*Network
}{}

func Start(address string) {

	// Setup keys
	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {

		fmt.Println("Generating keypair...")
		keypair = GenerateNewKeypair()
		WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}
	Core.Keypair = keypair

	// Setup Network
	Core.Network = SetupNetwork(address, BLOCKCHAIN_PORT)
	go Core.Network.Run()
	for _, n := range SEED_NODES() {
		Core.Network.ConnectionsQueue <- n
	}

	// Setup blockchain
	Core.Blockchain = SetupBlockchan()
	go Core.Blockchain.Run()

	go func() {
		for {
			select {
			case msg := <-Core.Network.IncomingMessages:
				HandleIncomingMessage(msg)
			}
		}
	}()
}

func CreateTransaction(txt string) *Transaction {

	t := NewTransaction(Core.Keypair.Public, nil, []byte(txt))
	t.Header.Nonce = t.GenerateNonce(TRANSACTION_POW)
	t.Signature = t.Sign(Core.Keypair)

	return t
}

func HandleIncomingMessage(msg Message) {

	switch msg.Identifier {
	case MESSAGE_SEND_TRANSACTION:
		t := new(Transaction)
		_, err := t.UnmarshalBinary(msg.Data)
		if err != nil {
			networkError(err)
			break
		}
		Core.Blockchain.TransactionsQueue <- t

	case MESSAGE_SEND_BLOCK:
		b := new(Block)
		err := b.UnmarshalBinary(msg.Data)
		if err != nil {
			networkError(err)
			break
		}
		Core.Blockchain.BlocksQueue <- *b
	}
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
