package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/izqui/blockchain/core"
)

var address = flag.String("ip", fmt.Sprintf("%s:%s", core.GetIpAddress()[0], core.BLOCKCHAIN_PORT), "Public facing ip address")

func init() {
	flag.Parse()
}

func main() {

	core.Start(*address)

	for {
		str := <-ReadStdin()
		core.Core.Blockchain.TransactionsQueue <- core.CreateTransaction(str)
	}
}

func ReadStdin() chan string {

	cb := make(chan string)
	sc := bufio.NewScanner(os.Stdin)

	go func() {
		if sc.Scan() {
			cb <- sc.Text()
		}
	}()

	return cb
}
