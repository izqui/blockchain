package main

import (
	"flag"
	"fmt"
	_ "github.com/izqui/helpers"
	"log"
	"time"
)

var (
	//flag
	address = flag.String("ip", GetIpAddress()[0], "Public facing ip address")

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

	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {

		fmt.Println("Generating keypair...")
		keypair = GenerateNewKeypair()
		WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}

	self.Keypair = keypair

	go RunBlockchainNetwork(*address, BLOCKCHAIN_PORT)

	time.Sleep(1000 * time.Second)
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
