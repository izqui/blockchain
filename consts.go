package main

const (
	BLOCKCHAIN_PORT      = "9119"
	MAX_NODE_CONNECTIONS = 1000

	NETWORK_KEY_SIZE = 80

	HEADER_SIZE = NETWORK_KEY_SIZE /* from key */ + NETWORK_KEY_SIZE /* to key */ + 4 /* int32 timestamp */ + 32 /* sha256 payload hash */ + 4 /* int32 payload length */

	KEY_POW_COMPLEXITY      = 0
	TEST_KEY_POW_COMPLEXITY = 0

	BLOCK_POW_COMPLEXITY      = 1
	TEST_BLOCK_POW_COMPLEXITY = 1

	KEY_SIZE = 28

	POW_PREFIX      = 0
	TEST_POW_PREFIX = 0
)
