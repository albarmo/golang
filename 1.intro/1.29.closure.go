package main

import "fmt"

func main(){
	fmt.Println("Closure || Block Scope")
	fmt.Println("=== Playground ===")	

	var counter int = 0
	name:= "Albar"
	increment := func(){
		// declare as new variable
		name:= "Moer"
		counter++

		fmt.Println(name)
	}
	
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

	fmt.Println("===================")
}