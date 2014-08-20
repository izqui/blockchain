package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type ConnectionsQueue chan string
type NodeChannel chan *Node
type Node struct {
	*net.TCPConn
	lastSeen int
}

type NodeSlice []*Node

func RunBlockchainNetwork() {

	listenCb := StartListening(*self.Address)

	fmt.Println("Listening in", *self.Address, BLOCKCHAIN_PORT)

	in, connectionCb := CreateConnectionsQueue()
	self.ConnectionsQueue = in

	for {
		select {
		case node := <-listenCb:
			fmt.Println("New connection from", node.TCPConn.RemoteAddr())
		case node := <-connectionCb:
			fmt.Println("Stablished connection to", node.TCPConn.RemoteAddr())
		}
	}
}

func CreateConnectionsQueue() (ConnectionsQueue, NodeChannel) {

	in := make(ConnectionsQueue)
	out := make(NodeChannel)

	go func() {

		for {
			address := <-in

			go ConnectToNode(address, false, out)
		}
	}()

	return in, out
}

func StartListening(address string) NodeChannel {

	cb := make(NodeChannel)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("[%s]:%s", address, BLOCKCHAIN_PORT))
	networkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	networkError(err)

	go func(l *net.TCPListener) {

		for {
			connection, err := l.AcceptTCP()
			networkError(err)

			cb <- &Node{connection, int(time.Now().Unix())}
		}

	}(listener)

	return cb
}

func ConnectToNode(dst string, retry bool, cb NodeChannel) {

	addrDst, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("[%s]:%s", dst, BLOCKCHAIN_PORT))
	networkError(err)

	var con *net.TCPConn = nil
	for {

		con, err = net.DialTCP("tcp", nil, addrDst)
		networkError(err)

		if con != nil {

			cb <- &Node{con, int(time.Now().Unix())}
			break
		} else {
			time.Sleep(5 * time.Second)
		}

		if !retry {
			break
		}
	}
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
