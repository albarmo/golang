package main

import "fmt"

func logging(){
	fmt.Println("Selesai memangil function..")
}

func runApplication (value int){
	// declare defer beginning of function
	defer logging()
	fmt.Println("Run Application")
	result:= 10/value
	fmt.Println(result)
}

func endApp(){
	// implement recover on defer function
	message := recover()
	if message != nil {
		fmt.Println("Error message : ",message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp2(error bool){
	defer endApp()
	if(error){
		panic("Aplikasi Error!")
	}
	fmt.Println("Aplikasi Berjalan...")
}


func main(){
	fmt.Println("Defer, Panic and Recover")
	fmt.Println("=== Playground ===")	
	fmt.Println("1. Always store recover on defer function")
	
	// defer
	runApplication(10)

	// panic and recover
	runApp2(true)


	fmt.Println("===================")
}