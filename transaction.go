package main

type TransactionSlice []Transaction
type Transaction struct {
	From      string
	To        string
	Data      string
	Timestamp int64
	Tries     int // nonce
}
