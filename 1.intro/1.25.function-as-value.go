package main

import "fmt"

func getGoodbye(name string) string{
	return "Goodbye " +  name 
}

func main (){
	fmt.Println("Function as Value")

	fmt.Println("=== Playground ===")	

	// store function in variable
	goodbye := getGoodbye
	fmt.Println(goodbye("Albar"))

	fmt.Println("===================")
}