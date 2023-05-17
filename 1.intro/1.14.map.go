package main

import "fmt"

func main (){
	fmt.Println("Map Data Type || Object")
	fmt.Println("1. Map is key-value")

	fmt.Println("=== Playground ===")

	var person map[string]string = map[string]string{
		"firstname": "Albar",
		"lastname": "Moerhamsa",
	}

	person["title"] = "Software Engineer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["firstname"])

	var book map[string]string = make(map[string]string)
	book["title"] =  "Belajar Albar"
	book["Year"] = "2023"
	book["Publisher"] = "Albar Cetakin"

	fmt.Println(book, len(book))
	delete(book, "Publisher")
	fmt.Println(book, len(book))
	fmt.Println("===================")
}