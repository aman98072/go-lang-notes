package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)
	if true {
		a := 2 // this is called variable shadowing
		a++
	}

	fmt.Println(a)
}
