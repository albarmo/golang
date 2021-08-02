package main

import "fmt"

func main(){
	fmt.Println("constant variable")

	const name = "Albar"

	const (	fistName string = "albar"
			lastName  = "moerhamsa"
		)

	fmt.Println(name)

	fmt.Println(fistName,"", lastName)
}