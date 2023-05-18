package main

import "fmt"

func main (){
	fmt.Println("Break and Continue In Golang")
	fmt.Println("=== Playground ===")
	for i:=0; i<= 10; i++{
		if i % 2 == 0{
			continue
		}else if i == 8 {
			break
		} 
		fmt.Println(i)
	}
	fmt.Println("===================")
}