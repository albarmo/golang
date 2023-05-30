package main

import "fmt"

type Man struct{
	Name string 
}

// using * for declare pointer on struct method
func (man *Man) Married (){
	man.Name ="Mr." + man.Name
	fmt.Println(man.Name,"Di Method")
}
func main (){
	fmt.Println("Pointer on Method | Struct")
	fmt.Println("=== Playground ===")
	/** Sangat direkomendasikan untuk menggunakan pointer di method,
	 untuk menghemat memory yang digunakan*/

	 var albar = Man{Name: "Albar"}
	 albar.Married()
	 fmt.Println(albar.Name)


	fmt.Println("===================")
}