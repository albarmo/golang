package main 

import "fmt"

func main(){
	fmt.Println("Array Data Types")
	fmt.Println("1. Data sotred on array must same")
	fmt.Println("2. Array has fixed lenght set when create array")
	fmt.Println("3. example : var array_name[length]data_type")
	fmt.Println("4. example create array with values : var array_name = [length]data_type{1,2,...}")

	fmt.Println("=== Playground ===")
	var names[2] string
	names[0] = "Albar"
	names[1] = "Moerhamsa"

	var fistname = names[0] 
	var lastname = names[1] 
	fmt.Println(names,fistname,lastname )

	var numbers = [10]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(numbers)
	fmt.Println(len(names), len(numbers))
	fmt.Println("===================")
}