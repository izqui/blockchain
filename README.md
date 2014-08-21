# Blockchain
Having fun implementing a blockchain using Golang.

Using [Minimum Viable Blockchain](https://artsec.hackpad.com/Blockchains-and-Bitcoins-mR2wlQ4KbVQ)

### Keys

The Blockchain uses ECDSA (224 bits) keys. 
When a user first joins the blockchain a random key will be generated.

Keys are encoded using base58.

Given x, y as the components of the public key, the key is generated as following:

```
	base58(BigInt(append(x as bytes, y as bytes)))
```

### Proof of work
In order to sign a transaction and send it to the network, proof of work is required. 

Proof of work is also required for block generation.

### Protocol

The blockchain runs on port `9191` and uses TCP to handle connections among peers.

##### Message

* Message type (1 byte)
	```
	const (
		MESSAGE_GET_NODES = iota + 20
		MESSAGE_SEND_NODES

		MESSAGE_GET_TRANSACTION
		MESSAGE_SEND_TRANSACTION

		MESSAGE_GET_BLOCK
		MESSAGE_SEND_BLOCK
	)
	```
* Options (4 bytes): Data specific
* Data (n bytes): Data specific

##### Transaction
	
* Header: 
	* From (80 bytes): Origin public key
	* To (80 bytes): Destination public key
	* Timestamp (4 bytes): int32 UNIX timestamp
 	* Payload Hash (32 bytes): sha256(payloadData)
	* Payload Length (4 bytes): len(payloadData)
	* Nonce (4 bytes): Proof of work

* Signature (80 bytes): signed(sha256(header))
* Payload data (Payload Length bytes): raw data

##### Block

* Header:
	* Origin (80 bytes): Origin public key
	* Timestamp (4 bytes): int32 UNIX timestamp
	* Previous block (32 bytes): sha256(previous block header)
	* Merkel Root (32 Bytes): sha256(transaction hashes)
	* Nonce (4 bytes): int32 UNIX timestamp

* Signature (80 bytes): signed(sha256(header))
* Block transactions
