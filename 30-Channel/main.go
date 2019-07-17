package main

import (
	"fmt"
)

func main() {
	hello := make(chan string, 2)
	hello <- "Hello"
	hello <- "Golang"
	close(hello)

	for v := range hello {
		fmt.Println(v)
	}
}
