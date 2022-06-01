package main // its mandatory defined for calling main function and main function must be call by compiler
// import fmt package for using print or println function
import "fmt"

// interface : its have object behaviours and structure contain object fields and must be defined all function of interface
type Employee interface {
	showEmpDetails()
	showEmpName() string
}

type Student struct {
	EmpId, EmpAge int
	EmpName       string
	EmpMobile     string
}

func main() {
	var obj Employee = Student{
		EmpId:     1,
		EmpAge:    23,
		EmpName:   "Aman",
		EmpMobile: "989898989",
	}

	obj.showEmpDetails()
	fmt.Println("Emp Name is : ", obj.showEmpName())
}

func (student Student) showEmpDetails() {
	fmt.Println("Emp id : ", student.EmpId)
	fmt.Println("Emp Age : ", student.EmpAge)
	fmt.Println("Emp Name : ", student.EmpName)
	fmt.Println("Emp Mobile No : ", student.EmpMobile)
}

func (student Student) showEmpName() string {
	return student.EmpName
}
