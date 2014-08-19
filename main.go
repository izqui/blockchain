package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	address = flag.String("ip", "::1", "Public facing ip address")
)

func main() {

	//GetIpAddress()
	keypair, _ := OpenConfiguration(HOME_DIRECTORY_CONFIG)
	if keypair == nil {
		//fmt.Println("Generating keypair...")
		keypair = GenerateNewKeypair(POW_PREFIX, KEY_POW_COMPLEXITY)
		//WriteConfiguration(HOME_DIRECTORY_CONFIG, keypair)
	}

	fmt.Println(len(keypair.Public))
	//fmt.Println(string(keypair.Private))
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
