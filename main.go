package main

import (
	"fmt"
	"log"
)

func main() {

	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {
		fmt.Println("Generating keypair")
		keypair = GenerateNewKeypair(POW_PREFIX, KEY_POW_COMPLEXITY)
		WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}

	fmt.Println(string(keypair.Public))
	fmt.Println(string(keypair.Private))
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
