package go_routines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	// ngubah thread
	runtime.GOMAXPROCS(20)
	// melihat jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println(totalGoRoutine)
}
