package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//const SEED_NODES = []string{"10.0.5.33"}

type BlockchainNode struct {
	*net.TCPConn
	lastSeen int
}

var nodes []BlockchainNode

func StartListening(address string) chan *BlockchainNode {

	cb := make(chan *BlockchainNode)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("[%s]:%s", address, BLOCKCHAIN_PORT))
	networkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	networkError(err)

	go func(l *net.TCPListener) {

		for {
			connection, err := l.AcceptTCP()
			networkError(err)

			cb <- &BlockchainNode{connection, int(time.Now().Unix())}
		}

	}(listener)

	return cb
}

func ConnectToNode(dst string, retry bool) chan *BlockchainNode {

	cb := make(chan *BlockchainNode)

	addrDst, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("[%s]:%s", dst, BLOCKCHAIN_PORT))
	networkError(err)

	go func() {
		var con *net.TCPConn = nil
		for {

			con, err = net.DialTCP("tcp", nil, addrDst)
			networkError(err)

			if !retry {
				break
			}

			if con != nil {
				cb <- &BlockchainNode{con, int(time.Now().Unix())}
				break
			} else {
				time.Sleep(5 * time.Second)
			}
		}

	}()

	return cb
}

func GetIpAddress() []string {

	name, err := os.Hostname()
	if err != nil {

		return nil
	}

	addrs, err := net.LookupHost(name)
	if err != nil {

		return nil
	}

	return addrs
}

func networkError(err error) {

	if err != nil {

		log.Println("Blockchain network: ", err)
	}
}
