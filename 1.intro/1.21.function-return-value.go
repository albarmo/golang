package main

import "fmt"

func perkalian(nilai1 int, nilai2 int) int{
	return nilai1 * nilai2
}

func getEmployeeSalary(position string) int{
	var salary int = 0
	switch position{
	case "Manager": 
		salary =  100000
	case "Lead": 
		salary = 80000
	case "Staff":
		salary =  600000
	}

	return salary
}

func main(){
	fmt.Println("Function - Return Value")
	fmt.Println("1. For function returned value must define data type of returned value in function")
	
	fmt.Println("=== Playground ===")	

	var resutl int = perkalian(8,10)
	fmt.Println(resutl)

	var salary int =  getEmployeeSalary("Manager")
	fmt.Println(salary)

	fmt.Println("===================")
}