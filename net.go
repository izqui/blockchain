package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/izqui/helpers"
)

type ConnectionsQueue chan string
type NodeChannel chan *Node
type Node struct {
	*net.TCPConn
	lastSeen int
}

type Nodes map[string]*Node

func (n Nodes) AddNode(node *Node) bool {

	key := node.TCPConn.RemoteAddr().String()

	if key != self.Address && n[key] == nil {

		fmt.Println("Node connected", key)
		n[key] = node
		return true
	}
	return false
}

func SetupNetwork(address, port string) ConnectionsQueue {

	in, connectionCb := CreateConnectionsQueue()
	self.Nodes = Nodes{}
	self.Address = fmt.Sprintf("%s:%s", address, port)

	fmt.Println("Listening in", self.Address)
	listenCb := StartListening(self.Address)

	go func() {
		for {
			select {
			case node := <-listenCb:

				self.Nodes.AddNode(node)
			case node := <-connectionCb:

				self.Nodes.AddNode(node)
			}
		}
	}()

	return in
}

func CreateConnectionsQueue() (ConnectionsQueue, NodeChannel) {

	in := make(ConnectionsQueue)
	out := make(NodeChannel)

	go func() {

		for {
			address := <-in

			address = fmt.Sprintf("%s:%s", address, BLOCKCHAIN_PORT)

			if address != self.Address && self.Nodes[address] == nil {

				go ConnectToNode(address, 5*time.Second, false, out)
			}
		}
	}()

	return in, out
}

func StartListening(address string) NodeChannel {

	cb := make(NodeChannel)
	addr, err := net.ResolveTCPAddr("tcp4", address)
	networkError(err)

	listener, err := net.ListenTCP("tcp4", addr)
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

func ConnectToNode(dst string, timeout time.Duration, retry bool, cb NodeChannel) {

	addrDst, err := net.ResolveTCPAddr("tcp4", dst)
	networkError(err)

	var con *net.TCPConn = nil
loop:
	for {
		breakChannel := make(chan bool)
		go func() {

			fmt.Println("Attempting to connect to", dst)
			con, err = net.DialTCP("tcp", nil, addrDst)

			if con != nil {

				cb <- &Node{con, int(time.Now().Unix())}
				breakChannel <- true
			}
		}()

		select {
		case <-helpers.Timeout(timeout):
			if !retry {
				break loop
			}
		case <-breakChannel:
			break loop
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
