package main

import "fmt"

func main (){
	fmt.Println("CONVERSION")
	fmt.Println("1. INTEGER CONVERSION")
	var nilai32 int32 = 200
	// var nilai64 int64 = int64(nilai32)
	var nilai8 int8 =  int8(nilai32)

	fmt.Println(nilai32,"Nilai 32")
	// fmt.Println(nilai64,"Nilai 64")

	//theres int overflow if the value cant reach minimun on int type
	fmt.Println("=== Integer overflow heree ====")
	fmt.Println(nilai8,"Nilai 8")

	fmt.Println("1. STRING CONVERSION")

	var name = "elang nichol"
	fmt.Println("len string buildin func =>", len(name))

	var pickedCharacter byte = name[11]
	var stringCharacter string = string(pickedCharacter)

	fmt.Println("Picked character", pickedCharacter)
	fmt.Println("Picked character after conversion", stringCharacter)
}