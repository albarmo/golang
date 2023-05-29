package main

import (
	"errors"
	"fmt"
)

func Pembagian (nilai int, pembagi int) (int,error){
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else{
		var result int = nilai/pembagi
		return result, nil
	}
}

func main (){
	fmt.Println("Error Interface")
	fmt.Println("=== Playground ===")

	// var errorExample  error =  errors.New("Upps Error!...")
	// fmt.Println(errorExample)

	hasil,err := Pembagian(100, 0)

	if err == nil {
		fmt.Println("Hasil : ", hasil)
	}else{
		fmt.Println(err)
	}
	
	fmt.Println("===================")
}