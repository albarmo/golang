package main 

import "fmt"

func sayHelloParams(name string, city string){
	fmt.Println("Halo dengan", name, "dari",city )
}

func main (){
	fmt.Println("Function - Parameters")
	fmt.Println("=== Playground ===")	
	sayHelloParams("Albar Moerhamsa","Jakarta")
	fmt.Println("===================")
}