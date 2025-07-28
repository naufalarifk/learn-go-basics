package belajar_golang

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("zamn")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(i int) {
	fmt.Println("Display", i)
}

func TestDisplayNumber(t *testing.T) {

	for i := range 100000 {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Microsecond)
}
