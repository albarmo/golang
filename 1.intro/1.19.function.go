package main

import "fmt"

func sayHello() {
	fmt.Println("Hellow!")
}

func love1000x(){
	for i:= 0; i<=1000; i++{
		fmt.Println(i," : I Love U")
	}
}

func main (){
	fmt.Println("Functions in Golang")
	fmt.Println("=== Playground ===")	
	sayHello()
	love1000x()
	fmt.Println("===================")
}