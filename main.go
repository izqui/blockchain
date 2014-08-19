package main

import (
	"fmt"
)

func main() {

	keypair := GenerateNewKeypair(POW_PREFIX, KEY_POW_COMPLEXITY)
	fmt.Println(string(keypair.Public))
	fmt.Println(string(keypair.Private))
}
