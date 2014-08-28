# The Blockchain

#### How to build a *distributed**, *immutable*, *tractable* data store
##### Jorge Izquierdo *-* Jonathan Dahan

--- 

# **Concepts** Transactions

![inline](https://1-ps.googleusercontent.com/s/www.igvita.com/posts/14/xtransaction-pki.png.pagespeed.ic.elb9fXIUMa.png)

---

# **Concepts** Transactions

```go
type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     uint32
	PayloadHash   []byte
	PayloadLength uint32
	Nonce         uint32
}
```

---


# **Concepts** Blocks

![inline](https://1-ps.googleusercontent.com/s/www.igvita.com/posts/14/xblockchain-full.png.pagespeed.ic.r5GP2Rwqya.png)

---

# **Concepts** Blocks

```go
type Block struct {
	*BlockHeader
	Signature []byte
	*TransactionSlice
}

type BlockHeader struct {
	Origin     []byte
	PrevBlock  []byte
	MerkelRoot []byte
	Timestamp  uint32
	Nonce      uint32
}
```

---

# **Concepts** Proof of work

```python




def proof_of_work(block, difficulty) {
	while (block.get_hash()[0:difficulty] != "0" * difficulty):
		block.header.nonce += 1
}
```

---

# **Properties** Decentralized

* Every peer has to download **all** the data.
* Every peer is connected over **TCP** to as many peers as possible.
* Every new transaction or block is **broadcast** to the network.

---

# **Properties** Inmutability

* It's really **hard to undo** because Proof of Work is expensive.
* Blocks refer to the hash of the previous block.
* The more blocks built on top of a block, the safer it is.


^ Undoing transaction 3 blocks down requires recalculating all of them

---

# Demo

---

# Satashi's gamble

* The cost of modifying a previously verified block must be **higher** than the benefit of verifying new blocks.

* We can now have an  append-only, **signed** database that can be completitely **decentralized**.

---

# What can you build on top?

**Bitcoin** peer-controlled currency
**Namecoin** distributed DNS service
**Certcoin?** web-of-trust style certificate store
**Keycoin?** distribution of public keys for any identity