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
	for _, n := range SEED_NODES {
		self.Network.ConnectionsQueue <- n
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
