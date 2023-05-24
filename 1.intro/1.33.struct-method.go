package main

import "fmt"


type Person struct{
	Name string
	Nationality string
	Age int
	IsMaried bool
}


func sayHi(customer Person){
	fmt.Println("Hi..,",customer.Name)
}

// struct method/function
func (customer Person) sayHiStructMethod(name string){
	fmt.Println("Hi..,",customer.Name, "Aku", name)
}

// struct method/function
func (customer Person) sayHuu(){
	fmt.Println("Huu..,",customer.Name)
}

func main(){
	fmt.Println("Struct Method")
	fmt.Println("1. Struct can store function")
	fmt.Println("=== Playground ===")

	person := Person{
		Name: "Albar",
		Nationality: "Indonesia",
		Age: 25,
		IsMaried: false,
	}
	// regular function 
	sayHi(person)

	person.sayHiStructMethod("Alyaa")
	person.sayHuu()
	
	fmt.Println("===================")
}