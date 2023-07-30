package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Albar"
	channel <- "Moerhamsa"
	channel <- "Bombastic"

	fmt.Println(<- channel,"1")
	fmt.Println(<- channel,"2")
	fmt.Println(<- channel,"3")

	time.Sleep(2 * time.Second)
	fmt.Println("=== FINISHED ===")
}