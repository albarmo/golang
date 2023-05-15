package main

import "fmt"

func main (){
	fmt.Println("Constant Variable in Golang")
	fmt.Println("1. Syntax constanta ==  const")
	fmt.Println("2. Data cant be changed")

	fmt.Println("=== Playground ===")
	const SECRET_CODE string = "AX19YZ"
	fmt.Println(SECRET_CODE)
	fmt.Println("===================")
}