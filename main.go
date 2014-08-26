package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	//flags
	address = flag.String("ip", GetIpAddress()[0], "Public facing ip address")

	self = struct {
		*Keypair
		*Blockchain
		*Network
	}{}
)

func init() {

	flag.Parse()
}
func main() {

	// Setup keys
	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {

		fmt.Println("Generating keypair...")
		keypair = GenerateNewKeypair()
		WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}
	self.Keypair = keypair

	// Setup Network
	self.Network = SetupNetwork(*address, BLOCKCHAIN_PORT)
	go self.Network.Run()
	for _, n := range SEED_NODES() {
		self.Network.ConnectionsQueue <- n
	}

	// Setup blockchain
	self.Blockchain = SetupBlockchan()
	go self.Blockchain.Run()

	// Read Stdin to create transactions
	stdin := ReadStdin()
	for {
		select {
		case str := <-stdin:
			self.Blockchain.TransactionsQueue <- CreateTransaction(str)
		case msg := <-self.Network.IncomingMessages:
			HandleIncomingMessage(msg)
		}
	}
}

func CreateTransaction(txt string) *Transaction {

	t := NewTransaction(self.Keypair.Public, nil, []byte(txt))
	t.Header.Nonce = t.GenerateNonce(TRANSACTION_POW)
	t.Signature = t.Sign(self.Keypair)

	return t
}

func HandleIncomingMessage(msg Message) {

	switch msg.Identifier {
	case MESSAGE_SEND_TRANSACTION:
		t := new(Transaction)
		err := t.UnmarshalBinary(msg.Data)
		if err != nil {
			networkError(err)
			break
		}
		self.Blockchain.TransactionsQueue <- t

	}
}

func ReadStdin() chan string {

	cb := make(chan string)
	sc := bufio.NewScanner(os.Stdin)

	go func() {
		for sc.Scan() {
			cb <- sc.Text()
		}
	}()

	return cb

}
func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
