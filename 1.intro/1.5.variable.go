package main

import "fmt"

func main(){
	fmt.Println("Variable in Golang")
	fmt.Println("1. One variable for one data type")
	fmt.Println("2. Declare variable using var (var varName int)")
	fmt.Println("3. Variable name is unique")
	fmt.Println("3. Multiple declaration variable")

	fmt.Println("=== Playground ===")
	var name string = "Albar Moerhamsa"
	fmt.Println(name)
	name = "Albar Wick"
	fmt.Println(name)
	
	var nickname =  "Baba Yaga"
	fmt.Println(nickname)

	var age int8 = 50
	var flag bool = false
	fmt.Println(age, flag)

	country := "Indonesia"
	fmt.Println(country,"< This variable declared using :=")

	var (
		firstName ="Albar"
		lastName = "Moerhamsa"
	)
	fmt.Println(firstName, lastName,"< Using multiple variable declaration")
	fmt.Println("===================")
}