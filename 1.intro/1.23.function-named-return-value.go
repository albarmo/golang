package main

import "fmt"

func getCompleteName()(firstName string, middleName string, lastName string){
	firstName = "Albar"
	middleName =  "Ganteng"
	lastName = "Moerhamsa"
	return 
}

func main (){
	fmt.Println("Function - Named Return Value")

	fmt.Println("=== Playground ===")	

	firstName, m, lastName := getCompleteName()
	fmt.Println(firstName, m, lastName)
	
	fmt.Println("===================")
}