package main // its mandatory defined for calling main function and main function must be call by compiler
// import fmt package for using print or println function
import "fmt"

func main() {
	// pointers
	var p = 34
	var p1 = &p
	fmt.Println(p, p1, *p1)
}
