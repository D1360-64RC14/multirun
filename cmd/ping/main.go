package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := "Dummy pinger"
	count := 0

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		fmt.Printf("%s: Pinging %d...\n", name, count)
		count++
		time.Sleep(1 * time.Second)
	}
}
