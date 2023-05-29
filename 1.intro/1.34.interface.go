package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHellow (name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName()string{
	return person.Name
}

type Animal struct{
	Name string
}

func (animal Animal) GetName()string{
	return animal.Name
}

func main (){
	fmt.Println("Interface in Golang")
	fmt.Println("=== Playground ===")
	/**
	Interface vs Struct
	- interface tidak memiliki implementasi langsung 
	- struct dapat memiliki implementasi secara langsung
	*/
	var customer Person
	customer.Name = "Albar"

	sayHellow(customer)
	animal := Animal{
		Name: "Vanila",
	}
	
	sayHellow(animal)

	fmt.Println("===================")
}