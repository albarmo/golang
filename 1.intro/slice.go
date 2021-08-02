package main

import "fmt"

func main (){
	fmt.Println("SLICEE")

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Spetember",
		"Oktober",
		"November",
		"Desember",
	}


	var slice1 = months[4:6]
	fmt.Println(months,"all months")
	fmt.Println(slice1," => slice 1")
	fmt.Println("length = ",len(slice1))
	fmt.Println("capacity = ", cap(slice1))
	slice1[0] = "ahgila"

	fmt.Println(months, "after reasing data on slice 1")

	slice2 := append(slice1, "AHHHHH")
	fmt.Println(slice2,"ini slice 2")

	slice3 := make([]string, len(slice2), cap(slice2))
	fmt.Println(slice3, "ini make slice dan sebelum di copy")

	copy(slice3, slice2)	

	fmt.Println(slice3, "after copy")
}