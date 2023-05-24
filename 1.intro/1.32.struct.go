package main

import "fmt"

func main () {
	fmt.Println("Struct")
	fmt.Println("=== Playground ===")
	
	type Customer struct {
		Name, Nationality string
		Age int
	}

	var person1 Customer
	person1.Name = "Albar"
	person1.Nationality= "INDONESIA"
	person1.Age = 25

	fmt.Println(person1)
	fmt.Println(person1.Name)

	/* STRUCT LITERALS*/
	// declare sturct data
	person2 := Customer{
		Name: "Alyaa",
		Nationality: "Uganda",
		Age: 24,
	}
	fmt.Println(person2)

	// its works but not recomended, it will breaks if theres changes in struct types
	person3 := Customer{"Joko","INDONEISA",40}
	fmt.Println(person3)

	fmt.Println("===================")
}