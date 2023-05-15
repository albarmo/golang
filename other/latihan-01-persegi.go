package main

import "fmt"

func main (){

	//type casting
	type result int8

	//declaration
	var sisi int = 5

	// mathematical operation
	var luasPersegi result =  result(sisi * sisi)
	var kelilingPersegi result = result(4 * sisi)

	// output
	fmt.Println("luas = ",luasPersegi)
	fmt.Println("keliling = ",kelilingPersegi)
}