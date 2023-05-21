package main

import "fmt"

// without type declaration
func sayHelloWithFilter(name string, filter func(string)string){
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

// With type declaration
type Filter func(string)string 
func sayHelloWithFilter2(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string{
	if name == "Anjing"{
		return "..."
	}else if name == "Fuckers" {
		return "..."
	}else{
		return name
	}
}


func main (){
	fmt.Println("Function as Parameter")

	fmt.Println("=== Playground ===")	

	filter := spamFilter
	sayHelloWithFilter("Boy", filter)
	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter2("Fuckers", filter)

	fmt.Println("===================")
}