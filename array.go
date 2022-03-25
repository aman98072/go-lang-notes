package main // its mandatory defined for calling main function and main function must be call by compiler
import "fmt" // import fmt package for using print or println function

func main() {
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
}
