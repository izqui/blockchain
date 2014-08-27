package main

import (
	"fmt"
	"time"

	"github.com/izqui/blockchain"
)

func main() {

	fmt.Println("Benchmarking...")
}

func benchmark(f func()) time.Duration {

	t0 := time.Now()

	f()

	return time.Now() - t0
}
