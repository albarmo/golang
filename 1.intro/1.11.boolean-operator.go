package main

import "fmt"

func main(){
	fmt.Println("Boolean Operation")
	fmt.Println("1. Has AND (&&) OR (||) NOT (!)")
	fmt.Println("=== Playground ===")

	var score int = 90
	const attendance int = 8

	type Result bool
	var isPassScore Result =  score >= 80
	var isPassAttendance Result =  attendance >=8

	var finalResult Result = isPassScore && isPassAttendance
	fmt.Println("FINAL RESULT : ",finalResult)
	fmt.Println(score >= 80 && attendance >= 8)
	
	fmt.Println("===================")
}