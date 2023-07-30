package go_routines

import (
	"fmt"
	"time"
	"testing"
)

func GiveMeResponse (channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Albar Moerhamsa"
}

func TestChannelAsParameter (t *testing.T) {
	channel :=  make(chan string)
	defer close(channel)
	
	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	
	time.Sleep(1 * time.Second)
}