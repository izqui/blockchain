package main

import (
	"fmt"
)

func main() {

	keypair := GenerateNewKeypair(POW_PREFIX, 1)
	fmt.Println(string(keypair.Public))
}
