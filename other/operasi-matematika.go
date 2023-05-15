package main

import "fmt"

func main (){
	var result = 10 + 8
	fmt.Println(result)

	// augmented assignment

	var a = 10
	a += 20
	
	fmt.Println(a)

	//unary operator
	var c =  10
	var checked bool = false
	c++ 
	fmt.Println(!checked)
	fmt.Println(c)
}