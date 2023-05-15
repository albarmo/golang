package main

import "fmt"

func main (){
	var nickname string
	nickname = "albarms"
	fmt.Println(nickname)

	var firstName = "Albar"
	var lastName = "Moerhamsa"
	var age = 24


	// shorthand to declare variable use :=

	country := "Indonesia"


	// declare multiple variables

	var (hobby = "Learn, Hiking"
	favoritePerson = "Alyaa Atiqoh")

	fmt.Println(nickname, ".Aka" , firstName, lastName, "age", age,"life in", country, "Hobby", hobby, "favorite person in the world" , favoritePerson)
}