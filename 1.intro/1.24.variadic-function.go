package main

import "fmt"

func sumAll(numbers ...int) int {
	var total int = 0
	for _, val := range numbers{
		total += val
	}
	return total
}

func countScore(name string, numbers ...int) (string, int) {
	var total int = 0
	for _, val := range numbers{
		total += val
	}
	return  name, total
}

func main (){
	fmt.Println("Variadic Function ")
	fmt.Println("1. Variadic function parameter must be on the end")

	fmt.Println("=== Playground ===")	
	result := sumAll(10,10,10,20,40,80)
	fmt.Println(result)

	name,result2 := countScore("albarms",10,20,30,40)
	fmt.Println(name,result2)
	
	// slice as input params
	scores := []int{10,20,30,40,50,60,70,80}
	result3 := sumAll(scores...)
	fmt.Println(result3)
	
	fmt.Println("===================")
}