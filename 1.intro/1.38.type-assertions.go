package main

import "fmt"

func random() interface{}{
	return "OK"
}

func main (){
	fmt.Println("Type Assertions")
	fmt.Println("=== Playground ===")
	/** 
	Merubah tipedata menajdi tipedata yang diinginkan 
	*/

	result:= random()
	resultString:=  result.(string)
	fmt.Println(resultString)

	// implement type assertion with switch
	switch value := result.(type){
	case string :
		fmt.Println("value", value,"is string")
	case int :
		fmt.Println("value", value,"is int")
	 default :
		fmt.Println("Uknown Type")
	}

	/** THIS SHOULD PANIC */
	// resultInt := result.(int)
	// fmt.Println(resultInt)
	
	fmt.Println("===================")
}