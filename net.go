package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//const SEED_NODES = []string{"10.0.5.33"}

type BlockchainNode struct {
	*net.TCPConn
	lastSeen int
}

var nodes []BlockchainNode

func StartListening(address string, port string) {

	addr, err := net.ResolveTCPAddr("tcp6", address)
	networkError(err)

	_, err = net.ListenTCP("tcp6", addr)
	networkError(err)
}

func GetIpAddress() {

	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	fmt.Println(addrs[0])
}

func networkError(err error) {

	if err != nil {

		log.Println("Blockchain network: ", err)
	}
}
