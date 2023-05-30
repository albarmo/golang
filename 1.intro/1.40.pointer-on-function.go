package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	fmt.Println(address)
	address.Country = "Indonesia"
}

func main (){
	fmt.Println("Pointer on Function")
	fmt.Println("=== Playground ===")

	var alamat Address = Address{"Jakarta", "DKI Jakarta", "X"}
	// without create variable pointer
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

	// with create variable pointer
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)


	fmt.Println("===================")
}