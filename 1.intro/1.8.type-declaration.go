package main

import "fmt"

func main (){
	fmt.Println("Type Declaration || Aliases")
	fmt.Println("1. syntax == type (type AliasesName data_types)")

	fmt.Println("=== Playground ===")
	type CardNumber string
	type Maried bool

	var cardId CardNumber = "012303131258"
	var mariedStatus Maried = false

	fmt.Println(cardId, mariedStatus)
	fmt.Println("===================")
}