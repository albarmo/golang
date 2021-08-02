package main

import "fmt"

func main(){
	// for declaration of new custom type 
	type nomorAnggota string
	var NIK nomorAnggota = "HP-146"

	fmt.Println(NIK)
}