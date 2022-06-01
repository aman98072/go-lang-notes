package main // its mandatory defined for calling main function and main function must be call by compiler

import "fmt" // import fmt package for using print or println function

var Myname = "Aman"   // global variable
var myLName = "Gupta" // package level variable or myname also package level when define in global

// structure
type Employee struct {
	EmpId     int
	EmpName   string
	EmpMobile []string
}

func main() {
	fmt.Println("Hello world") // println("hello world") // for print a msg

	// define diff type of variables
	var a = 10 // local variable
	var b int = 34
	c := 23
	var h = 32
	var i = h
	const uname = "Yes"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(Myname)
	fmt.Println(myLName)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(uname)
	fmt.Printf("%v %T", h, h)
	println()

	// Array
	var arr = [...]int{1, 2, 4, 5}
	fmt.Println(arr, arr[0], len(arr)) // len for getting length of array

	// slice from array
	fmt.Println(arr[1:3]) // 1 is staring and 3 is no of elements
	fmt.Println(arr[:2])  // 2 means no of element till
	fmt.Println(arr[2:])  // 2 means no of element left

	// two demional Array
	var matrix = [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(matrix, matrix[0][1])

	// array slicing
	var sarr = []int{1, 2, 4, 5}
	var sarr1 = sarr
	sarr[2] = 34
	fmt.Println("i am slicing array : ", sarr)
	fmt.Println("i am copy of slicing array : ", sarr1) // now here slicing array just copy address not values
	fmt.Println("Check capacity of array : ", cap(sarr))

	// append values in array
	var newArr = []int{1, 3, 4, 6}
	newArr = append(newArr, 45, 34)
	fmt.Println("Appended array values are : ", newArr)

	// map
	var emp = map[string]int{
		"Aman":  11200,
		"Rahul": 3000,
		"pawan": 3400,
	}

	// check key is exist or not
	var empKeys, checkKey = emp["Neha"]
	// var _, checkKey = emp["Neha"] // here _ means will be use _ variable in future then its not through the error
	fmt.Println("Check key is exists or not : ", empKeys, checkKey) // here false is define for not define key in emp its return boolean value

	var emp1 = emp // when assign variable in map its using address of direct variable not creating own address
	emp["pankaj"] = 45

	fmt.Println(emp)
	fmt.Println(emp1)
	delete(emp, "pankaj") // delete key in map
	fmt.Println(emp)
	fmt.Println(emp1)

	//structure
	var employee = Employee{
		EmpId:     1,
		EmpName:   "Aman",
		EmpMobile: []string{"8986545645", "534534545s"},
	}

	employee.EmpId = 2
	var employee1 = &employee // here u can use reference of employee
	fmt.Println("employee struct : ", employee.EmpId, employee.EmpMobile[0], employee1)

	// if else
	var j = 4
	if j == 5 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

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

	// loops
	// there is only one type loop in go lang
	for i := 1; i <= 5; i++ {
		fmt.Println("hello world", i)
	}

	//defer : its behave like stack means first in last out its basically used for after excute the code like in db connection closing
	defer fmt.Println("DB closing")
	defer fmt.Println("manipulate DB query")
	defer fmt.Println("Connection create")

	// pointers
	var p = 34
	var p1 = &p
	fmt.Println(p, p1, *p1)

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
