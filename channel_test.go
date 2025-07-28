package belajar_golang

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Naufal Arif Kurniawan"
		fmt.Println("Done!")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveMeRes(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Naufal Arif Kurniawan"
}

func TestGimmeRes(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRes(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

//channel chan <- string indicates this params only receives channel

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "Naufal Arif Kurniawannn"
}

//channel chan <- string indicates this params only sends channel

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "NaU"
		channel <- "NaU"
		channel <- "NaU"
		channel <- "NaU"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Done!")
}

func TestRangeChan(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := range 10 {
			channel <- "Iteration No. " + strconv.Itoa(i)
		}
		defer close(channel)

	}()

	for data := range channel {
		fmt.Println("Receiving Data", data)

	}
	fmt.Println("Done!")

}
