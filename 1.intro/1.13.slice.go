package main

import "fmt"

func main(){
	fmt.Println("Slice Data Type")
	fmt.Println("1. Slice adalah potongan dari array")
	fmt.Println("2. Ukuran data slice bisa berubah")
	fmt.Println("3. Slice memiliki pointer, length and capacity")
	fmt.Println("4. Slice is dynamic array is fixed")

	fmt.Println("=== Playground ===")
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

	slice1 := months[6:9]
	slice1[0] = "MayX"
	fmt.Println(months, len(months))
	fmt.Println("Slice 1 : ", slice1,"Length : ", len(slice1),"Capacity : ", cap(slice1))
	

	var days = [7]string{"Senin", "Selasa", "Rabu","Kamis","Jumat","Sabtu","Minggu"}
	daySlice := days[5:]
	//append
	daySlice2 := append(daySlice,"Albar")
	
	fmt.Println(days)
	fmt.Println(daySlice)
	fmt.Println(daySlice2)

	//make
	newSlice := make([]string, 2, 5)
	newSlice[0] =  "Albar"
	newSlice[1] = "Moerhamsa"

	fromSlice := newSlice[:]
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, fromSlice)

	// if lenght not declared on array it will be slice
	// ex: array (var array[10]string) 
	// ex: slice (var slice[]string)

	fmt.Println(fromSlice ," copied >>> ", copySlice)
	fmt.Println("===================")
}