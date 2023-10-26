package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterSlow(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // slow simulation
			}
		}
	}()
	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total GoRoutine ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterSlow(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total GoRoutine ", runtime.NumGoroutine())
}
