package main // its mandatory defined for calling main function and main function must be call by compiler
import "fmt" // import fmt package for using print or println function

// structure its a group of collection of variables
type Employee struct {
	EmpId     int
	EmpName   string
	EmpMobile []string
}

func main() {
	//structure
	var employee = Employee{
		EmpId:     1,
		EmpName:   "Aman",
		EmpMobile: []string{"8986545645", "534534545s"},
	}

	employee.EmpId = 2
	var employee1 = &employee // here u can use reference of employee
	fmt.Println("employee struct : ", employee.EmpId, employee.EmpMobile[0], employee1)
}
