package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCond(val int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()

	fmt.Println(val)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := range 10 {
		go WaitCond(i)
	}

	// go func() {
	// 	for range 10 {
	// 		time.Sleep(1 * time.Millisecond)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
