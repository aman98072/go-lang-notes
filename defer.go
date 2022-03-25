package main // its mandatory defined for calling main function and main function must be call by compiler
// import fmt package for using print or println function
import "fmt"

func main() {
	//defer : its behave like stack means first in last out its basically used for after excute the code like in db connection closing
	defer fmt.Println("DB closing")
	defer fmt.Println("manipulate DB query")
	defer fmt.Println("Connection create")
}
