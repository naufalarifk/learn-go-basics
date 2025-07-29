package belajar_golang

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for range 1000 {
		go func() {
			group.Add(1)
			for range 100 {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	assert.Equal(t, int64(100000), x)
	fmt.Println("Counter", x)

}
