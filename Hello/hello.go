package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(hello())
}

func hello() string {

	return quote.Hello()
}
