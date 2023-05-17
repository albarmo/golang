package main

import "fmt"

func main(){
	fmt.Println("If Expression on Golang")
	fmt.Println("1. If have a Short Statement")
	fmt.Println("=== Playground ===")
	var person map[string]string = make(map[string]string)
	person["name"] =  "Alyaax"
	person["hobby"] = "Ngunyah"
	fmt.Println(person)

	if person["name"] == "Alyaa Atiqoh"{
		fmt.Println("Ngunyah ngunyah ngunyah!!!")
	} else if person["hobby"] ==  "Ngunyah" {
		fmt.Println("Alya ini pasti")
	}else{
		fmt.Println("Bukan nih, yakali ga ngunyah")
	}

	// Short Statement
	if lenght := len(person["name"]); lenght >= 10{
		fmt.Println(lenght)
		fmt.Println("Terlalu Panjang!")
	}else{
		fmt.Println("Sudah valid!")
	}

	fmt.Println("===================")
}