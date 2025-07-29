package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C

	fmt.Println(time)

}

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)

	fmt.Println(time.Now())
	time := <-channel

	fmt.Println(time)

}

func TestAfterFunc(t *testing.T) {

	group := sync.WaitGroup{}

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}

// below has deadlock. why? channel isnt closed and all channel cannot fit in time ticker. use select instead

func TestTickerDeadlock(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}

}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	// timeval using select allows listening to multiple channels
	for {
		select {
		case timeVal := <-ticker.C:
			fmt.Println(timeVal)
		case <-done:
			return
		}
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}

}
