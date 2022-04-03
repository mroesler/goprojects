package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in the FOO function")
}

// control flow:
// (1) sequence
// (2) loop
// (3) conditional
