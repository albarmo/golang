package main

import "fmt"

func main(){
	fmt.Println("For Loop in Golang")
	fmt.Println("1. Go only have for loops")

	fmt.Println("=== Playground ===")
	counter:= 1

	// Unstatement Loops
	for counter <= 10 {
		fmt.Println("Iterasi ke : ", counter)
		counter++
	}

	// Statement Loops
	for i:= 0; i<= 10; i++{
		fmt.Println("Iterasi ke : ", i)
	}

	var persons []string = []string{"Albar","Alyaa","Eki"}
	for j:= 0; j < len(persons); j++{
		fmt.Println(persons[j])
	}	
	
	// For Range 
	for index, item:= range persons{
		fmt.Println(index ," : ", item)
	}

	// if index not needed can replace with _
	for _, value := range persons{
		fmt.Println(value)
	}


	// Mapping Map || Object
	logsActivity := make(map[string]string)
	logsActivity["id"] = "0"
	logsActivity["title"] = "Belajar Go Lang"

	for key, activity := range logsActivity{
		fmt.Println(key, " : ", activity)
	}

	fmt.Println("===================")
}