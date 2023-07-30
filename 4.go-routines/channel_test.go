package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	// // declare a channel
	// channel := make(chan string)

	// // send data to channel
	// channel <- "Some string data"

	// // recive data from channel
	// var data string = <- channel
	// fmt.Println("Stored in Variable : ", data) 

	// // direct use as parameter
	// fmt.Println("Channel Data : ", <- channel)
	// close(channel)


	channelTwo  := make(chan string) 
	defer close (channelTwo)
	go func (){
		time.Sleep(1 * time.Second)
		channelTwo <- "Albar Moerhamsa"
		fmt.Println("Selesai megirim data ke channel")
	}()

	var data string  = <- channelTwo
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	
	// #is recomended to close channel after its done
}