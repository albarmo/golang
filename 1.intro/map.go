package main 


import "fmt"

func main (){

	var person = map[string]string{
		"name" : "Albar",
		"address" : "Boulevard of broken dreams",
	}

	fmt.Println(person)
	person["age"] =  "20"
	fmt.Println(person)
	fmt.Println("map length = ",len(person))


	// var book = make(map[string]string)
	//or
	//with specifict data typee
	var book map[string]string = make(map[string]string)
	book["title"] = "Gedebook"
	book["author"] = "Made by mistake"
	book["price"] = "20.000"
	book["data"] = "xxx"

	fmt.Println(book,"ini book setelah diisi")
	fmt.Println("book map length = " , len(book))

	//delete key on map 
	delete(book, "data")
	
	fmt.Println(book,"ini book setelah diisi")
	fmt.Println("book map length = " , len(book))






}