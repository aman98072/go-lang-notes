package main // its mandatory defined for calling main function and main function must be call by compiler
// import fmt package for using print or println function
import "fmt"

// structure
type Employee struct {
	EmpId     int
	EmpName   string
	EmpMobile []string
}

func main() {
	// functions
	hello()
	fmt.Println("Sum of two no is : ", add(4, 5))
	fmt.Println("Sum of many no is : ", addAll(4, 5, 6, 1))
	var sum, sub, mul, div = calc(20, 10)
	fmt.Println("Sum of two no is : ", sum)
	fmt.Println("Sub of two no is : ", sub)
	fmt.Println("Mul of two no is : ", mul)
	fmt.Println("Div of two no is : ", div)

	// function with structure :
	var employee = Employee{
		EmpId:     1,
		EmpName:   "Aman",
		EmpMobile: []string{"8986545645", "534534545s"},
	}
	employee.showEmpInfo()

	// short hand function
	f := func() string {
		return "Hello i am anonymous function"
	}

	fmt.Println(f())

	// it self calling function IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("Hi i am IIFE function (Immediately Invoked Function Expression)")
	}()
}

func hello() {
	fmt.Println("Here i am function : Hello world")
}

func add(a, b int) int { // syntax
	return a + b
}

func addAll(vals ...int) int { // variadic function its like spread operator in js
	var sum = 0
	for _, v := range vals { // its like foreach loop
		sum += v
	}

	return sum
}

func calc(a int, b int) (int, int, int, int) { // return multiple types of datatype
	return (a + b), (a - b), (a * b), (a / b)
}

func (emp Employee) showEmpInfo() {
	fmt.Println("Emp id : ", emp.EmpId)
	fmt.Println("Emp Name : ", emp.EmpName)
	fmt.Println("Emp Mobile Number : ", emp.EmpMobile)
}
