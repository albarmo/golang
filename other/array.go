package main

import "fmt"

func main () {
	fmt.Println("ARRAY DI GO")
	
	var name[4] string
	name[0] = "Albar"
	name[1] = "Moerhamsa"
	name[2] = "Alyaa"
	name[3] = "Atiqooh"
	
	fmt.Println(name[0],"&", name[2])
	fmt.Println(len(name))
	
	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	// function len itu panjang arraynya bukan jumlah datanya
}