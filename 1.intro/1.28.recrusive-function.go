package main

import "fmt"

// using for loop functon
func factorialLoop (value int)int{
	result := 1
	for i:= value; i > 0; i--{
		result *= i
	}
	return result
}

// using recrusive function
func factorialRecrusiveFunction(value int) int{
	if(value == 1){
		return 1
	}else{
		return value * factorialRecrusiveFunction((value-1)) 
	}
}

func main(){
	fmt.Println("Recrusive Function")
	fmt.Println("=== Playground ===")	
	
	loop :=  factorialLoop(10)
	recrusive:= factorialRecrusiveFunction((10))
	fmt.Println("Loop : ",loop)
	fmt.Println("Recrusive : ",recrusive)

	fmt.Println("===================")
}