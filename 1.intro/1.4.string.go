package main

import "fmt"

func main (){
	fmt.Println("String Data Types")
	fmt.Println("1. String represented with string or double petik")
	fmt.Println("2. String has function len()")
	fmt.Println("3. String can be accesed using string_variable[index]")
	fmt.Println("It return character bytes using string() to convert to char")


	fmt.Println("=== Playground ===")
	fmt.Println("Albar")
	fmt.Println("Albar Moerhamsa")
	fmt.Println(len("Albar"))
	fmt.Println(("Albar"[0]))
	fmt.Println("Albar Moerhamsa"[1])

	var name string= "Albar Moerhamsa"
	var char string = string(name[0])
	fmt.Println(char)
	fmt.Println("===================")
}