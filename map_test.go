package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, sync *sync.WaitGroup) {
	defer sync.Done()
	sync.Add(1)

	data.Store(value, value)

}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := range 100 {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)

		return true
	})
}

//pro tip: instead of golang's Map, use sync's Map. it's safer
