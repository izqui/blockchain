package main

import (
	"fmt"
	"github.com/izqui/helpers"
	"reflect"
	"time"
)

const (
	TRANSACTION_COMPLEXITY = 2
)

func (t *Transaction) Init() []byte {

	prefix := helpers.ArrayOfBytes(TRANSACTION_COMPLEXITY, 0) // [0 0]
	for !reflect.DeepEqual(t.Hash()[0:TRANSACTION_COMPLEXITY], prefix) {

		t.Tries += 1
	}

	return t.Hash()
}
func (t *Transaction) Hash() []byte {

	txt := fmt.Sprintf("%s%s%s%s%s", t.From, t.To, t.Data, t.Timestamp, t.Tries)
	return helpers.SHA256([]byte(txt))
}

func main() {

	t := Transaction{From: "Jorge", To: "Jonathan", Data: "Hello", Timestamp: time.Now().Unix()}
	fmt.Println(t)
	fmt.Printf("%x\n", t.Init())
	fmt.Println(t.Hash())
	fmt.Println(t.Tries)
}
