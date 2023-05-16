package main

import "fmt"

func main(){
	fmt.Println("Comparation Operator in Golang")

	fmt.Println("=== Playground ===")
	var (
	a = 21
	b = 20
	)
	var isGreaterThan bool =  a > b 
	fmt.Println(a,">", b," : ", isGreaterThan)

	type Nama string
	var peopleOne Nama = "Albar" 
	const PeopleTwo Nama = "ALbar"

	var isSamePerson bool =  peopleOne > PeopleTwo 
	fmt.Println(peopleOne,">", PeopleTwo," : ",isSamePerson)
	fmt.Println("===================")
}