package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/izqui/helpers"
)

var (
	//flag
	address = flag.String("ip", GetIpAddress()[0], "Public facing ip address")

	// TODO: Reduce to Keypair, Blockchain, Network [Issue: https://github.com/izqui/blockchain/issues/1]
	self = struct {
		*Keypair
		*Blockchain
		Nodes
		ConnectionsQueue
		Address string
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
	self.ConnectionsQueue = SetupNetwork(*address, BLOCKCHAIN_PORT)
	for _, n := range SEED_NODES {
		self.ConnectionsQueue <- n
	}

	// Setup blockchain
	self.Blockchain = SetupBlockchan()
	go self.Blockchain.Run()

	// Read Stdin to create transactions
	stdin := ReadStdin()
	for {
		st := <-stdin

		t := NewTransaction(keypair.Public, nil, []byte(st))
		t.Header.Nonce = t.GenerateNonce(helpers.ArrayOfBytes(TRANSACTION_POW_COMPLEXITY, POW_PREFIX))
		t.Signature = t.Sign(keypair)

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
