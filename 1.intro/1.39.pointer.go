package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main (){
	fmt.Println("Pointer")
	fmt.Println("=== Playground ===")
	fmt.Println("1. In Golang using Pass By Value")
	fmt.Println("2. Pointer used for Pass By Reference")
	fmt.Println("3. Operator ( & * ) Pointer")

	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 is a pointer
	var address2 *Address = &address1
	var address3 *Address = &address1
	// create data with empty
	var address4 *Address = new(Address)

	address2.City = "Bandung"
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)


	fmt.Println("===================")
}