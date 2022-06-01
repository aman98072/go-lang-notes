package main

import (
	"fmt"
)

func main() {
	var val interface{} = 34
	var val1 int = val.(int) // this is type assertion
	// var val1 int = val.(string) // throw error becuase type not match
	fmt.Println(val, val1)
}
