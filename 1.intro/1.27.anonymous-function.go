package main

import "fmt"


type BlackList func(string)bool

func registerUser(name string, blackList BlackList){	
	if(blackList(name)){
		fmt.Println("You are blocked!", name)
	}else{
		fmt.Println("Welcomee!",name)
	}
}

func main (){
	fmt.Println("Function as Value")

	fmt.Println("=== Playground ===")	
	blacklist := func(name string)bool{
		return name == "Admin"
	}


	registerUser("Admin", blacklist)
	registerUser("Albar", blacklist)	
	registerUser("root", func(name string)bool{
		return name == "root"
	})	

	fmt.Println("===================")
}