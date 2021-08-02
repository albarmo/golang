package main

import "fmt"


func main () {
	var nilai int =  90
	var absen int = 9

	var lulusUjian = nilai > 90
	var lulusAbsen =  absen >=  9

	var lulus bool =  lulusUjian && lulusAbsen
	var nyogok bool = lulusAbsen || lulusUjian
	var bapaknyaYangPunyaKampus bool = !lulus
	var ngecit bool = (nilai + 10) > 90 && (absen + 1) > 9

	fmt.Println(lulus)
	fmt.Println(nyogok)
	fmt.Println(bapaknyaYangPunyaKampus)
	fmt.Println(ngecit)

}