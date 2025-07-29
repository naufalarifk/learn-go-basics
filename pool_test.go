package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	Wait := &sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Naufal")
	pool.Put("Arif")
	pool.Put("Kurniawan")

	for range 10 {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			Wait.Wait()
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Done!")

}
