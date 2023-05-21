package main

import "fmt"

// define all returning values
func getFullName()(string,string, int){
	return "Albar", "Moerhamsa", 25
}

func main (){
	fmt.Println("Function -  Return Multiple Value")

	fmt.Println("=== Playground ===")	

	// calling and store value in variable
	// use _ to ignore returning variable
	firstName,lastName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	
	fmt.Println("===================")
}