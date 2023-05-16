package main

import "fmt"

func main(){
	fmt.Println("Math Operation")

	fmt.Println("=== Playground ===")
	type Numberical int
	var a int = 10
	var b int = 10

	c :=  a + b
	fmt.Println("result : ",c)

	c+= 10
	fmt.Println("result : ",c)

	c++ 
	fmt.Println("result : ",c)

	var d = -c 
	fmt.Println("result : ",d)
	fmt.Println("===================")
}
