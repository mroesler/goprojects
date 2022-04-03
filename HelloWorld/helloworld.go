package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in the FOO function")
}

func bar() {
	var myAge int = 41
	myAge++
	fmt.Printf("%s%d%s", "Next year i will be ", myAge, " years old.\n")
}

// control flow:
// (1) sequence
// (2) loop
// (3) conditional
