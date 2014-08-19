package main

import (
	"net"
)

//const SEED_NODES = []string{"10.0.5.33"}

type BlockchainNode struct {
	*net.TCPConn
	lastSeen int
}

var nodes []BlockchainNode
