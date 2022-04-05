package main

import (
	"fmt"

	"rsc.io/quote"
)

type MICHA int

var x MICHA

//var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//hello()
	//assignVariables()
	// fmt.Println(x, y, z)
	//concatXYZ()

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

func hello() {
	fmt.Println(quote.Hello())
}

func assignVariables() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%d\n%s\n%t\n", x, y, z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}

func concatXYZ() {
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
}
