package main // its mandatory defined for calling main function and main function must be call by compiler
import "fmt" // import fmt package for using print or println function

func main() {
	// switch case
	// in switch case donot need break keyword its only run when condition match
	// fallthrough : when condition true then below one only case will run automatically
	var k = 1
	switch k {
	case 1:
		fmt.Println("First")
		fallthrough
	case 2:
		fmt.Println("Second")
	case 3, 5:
		fmt.Println("Third")
	case 4, 8: // here u can use multiple case with comma seprated value
		fmt.Println("Fourth")
	default:
		fmt.Println("Wrong inputs")
	}
}
