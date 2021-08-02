package main

import "fmt"

func main (){
	var nickname string
	nickname = "albarms"
	fmt.Println(nickname)

	var age = 23
	fmt.Println(age)
	
	//  this line will error because variable age was declared with int value
	// age = "dua puluh tiga tahun"
	// fmt.Println(age)

	firstName := "Albar"
	lastName := "Moerhamsa"
	fmt.Println(firstName, lastName)

	var(
		shesFirstName string =  "Alyaa"
		shesLastName = "Atiqoh"
	)

	fmt.Println(shesFirstName + shesLastName)

}