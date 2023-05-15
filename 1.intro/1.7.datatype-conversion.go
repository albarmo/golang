package main

import "fmt"

func main (){
	fmt.Println("Constant Variable in Golang")
	fmt.Println("1. Beware integer overflow when converting")

	fmt.Println("=== Playground ===")
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 =  int8(nilai64)

	fmt.Println(nilai32, nilai64, nilai8)
	fmt.Println("===================")
}