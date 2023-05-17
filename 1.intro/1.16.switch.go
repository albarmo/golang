package main

import "fmt"

func main (){
	fmt.Println("Switch in Golang")
	fmt.Println("1. Switch have a Short Statement")

	fmt.Println("=== Playground ===")
	var name string = "Albar"

	switch name {
	case "Albar":
		fmt.Println("Halo Albar")
	case "Alyaa":	
		fmt.Println("Halo Alyaa")
	default:
		fmt.Println("Halo Semuanya!!!")
	}

	switch length := len(name); length > 10{
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false : 
		fmt.Println("Sudah Benar")
	}

	firstCharacter := string(name[0])

	fmt.Println(firstCharacter)
	switch{
	case firstCharacter == "A":
		fmt.Println("Alyaa")
	case firstCharacter  == "B":
		fmt.Println("Bagas")
	case firstCharacter == "C":
		fmt.Println("Cecep")
	default:
		fmt.Println("Mantab")
	}

	fmt.Println("===================")
}