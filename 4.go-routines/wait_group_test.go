package go_routines

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func RunAsyncronous(group *sync.WaitGroup, numb int){
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello", numb)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group, i)
	}

	group.Wait()
	fmt.Println("Complete!")
}