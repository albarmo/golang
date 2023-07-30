package go_routines

import (
	"fmt"
	"testing"
	"sync"
	"time"
)

func TestPool(t *testing.T) {
	// var pool = sync.Pool
	// Wiht default value on pool
	 pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	
	pool.Put("Albar")
	pool.Put("Alyaa")
	pool.Put("Moerhamsa")
	pool.Put("Atiqoh")

	for i := 0; i < 10; i++ {
		go func(){
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)

		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")

}

