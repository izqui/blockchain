package main

import (
	"flag"
	"fmt"
	_ "github.com/izqui/helpers"
	"log"
)

var (
	address = flag.String("ip", GetIpAddress()[0], "Public facing ip address")
)

func init() {

	flag.Parse()
}
func main() {

	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {

		fmt.Println("Generating keypair...")
		keypair = GenerateNewKeypair(POW_PREFIX, KEY_POW_COMPLEXITY)
		WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}

	listenCb := StartListening(*address)
	fmt.Println("Listening in", *address, BLOCKCHAIN_PORT)

	connectCb := ConnectToNode(*address, true)
	for {
		select {
		case node := <-listenCb:
			fmt.Println("New connection from", node.TCPConn.RemoteAddr())
		case node := <-connectCb:
			fmt.Println("Connected to", node.TCPConn.RemoteAddr())
		}

	}
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
