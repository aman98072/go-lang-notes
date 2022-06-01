package main // its mandatory defined for calling main function and main function must be call by compiler
import "fmt" // import fmt package for using print or println function

func main() {
	// loops
	// there is only one type loop in go lang which is for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("hello world", i)
	}

	// short hand like foreach no need to pass condition
	var sum = 0
	var vals = [3]int{2, 4, 5}
	for _, v := range vals { // its like foreach loop
		sum += v
	}
	fmt.Println("Total sum is : ", sum)

	// run loop as infinate
	for {
		fmt.Println("hello")
	}
}
