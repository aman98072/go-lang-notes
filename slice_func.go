package main

import "fmt"

func main() {
	a := make([]int, 2, 5)            // its slice function
	fmt.Println(a)                    // [0 0]
	a = append(a, 4, 6, 7, 3, 34, 65) // append multiples values
	fmt.Println(a)                    // [0 0 4 6 7 3 34 65]
}
