package go_routines

import (
	"fmt"
	"time"
	"testing"
)

func GiveMeResponse2 (channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Albar Moerhamsa"
}

func TestChannelSelect (t *testing.T){
	channelOne := make(chan string)
	channelTwo := make(chan string)
	defer close(channelOne)
	defer close(channelTwo)

	go GiveMeResponse2(channelOne)
	go GiveMeResponse2(channelTwo)

	var counter int = 0
	for {
		select {
		case data := <- channelOne:
			fmt.Println("Data dari channel satu : ", data)
			counter++
		case data := <- channelTwo:
			fmt.Println("Data dari channel satu : ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestChannelDefaultSelect (t *testing.T){
	channelOne := make(chan string)
	channelTwo := make(chan string)
	defer close(channelOne)
	defer close(channelTwo)

	go GiveMeResponse2(channelOne)
	go GiveMeResponse2(channelTwo)

	var counter int = 0
	for {
		select {
		case data := <- channelOne:
			fmt.Println("Data dari channel satu : ", data)
			counter++
		case data := <- channelTwo:
			fmt.Println("Data dari channel satu : ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}