package belajar_golang

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	group := sync.WaitGroup{}
	for range 100 {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total GO Routine", totalGoRoutine)

	group.Wait()

}
