package go_routines

import (
	"fmt"
	"testing"
	"strconv"
)

func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	
	for data := range channel{
		fmt.Println("Menerima data", data)
	}

	fmt.Println("=== FINISHED ===")
}