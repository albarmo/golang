package main

import "fmt"

func main () {
	fmt.Println("OPERATOR PERBANDINGAN")


	var name1 = "ALBAR"
	var name2 = "ALBAR"

	fmt.Println(name1 > name2)

	var value1 int = 10
	var value2 int = 10

	type returnType bool
	var result returnType = value2 > value1

	fmt.Println("RESULT : " , result )

	fmt.Println(value1 ==  value2)
	fmt.Println(value1 !=  value2)
	fmt.Println(value1 >=  value2)
	fmt.Println(value1 <=  value2)
	fmt.Println(value1 >  value2)
	fmt.Println(value1 <  value2)

	// oeprator : >= <= ! != 

}
