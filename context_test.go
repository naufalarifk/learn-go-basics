package belajar_golang

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	bg := context.Background()

	fmt.Println(bg)

	todo := context.TODO()

	fmt.Println(todo)
}

func TestContextWithVal(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	ctxF := context.WithValue(ctxC, "f", "F")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)

	fmt.Println(ctxA.Value("b")) // parent can't get val from child
	fmt.Println(ctxF.Value("c")) // child can get val from parent
	fmt.Println(ctxF.Value("b")) // different parent, can't get val
	fmt.Println(ctxF.Value("f")) // can get its own data (obviously)
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return

			case destination <- counter:
				counter++
				time.Sleep(1 * time.Second) //simulate slow process
				// default:
				// destination <- counter:
				// 	counter++
			}
		}

	}()

	return destination

}

func TestCtxWithCancel(t *testing.T) {
	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())

}

func TestCtxWithTimeout(t *testing.T) {
	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())

}

func TestCtxWithDeadline(t *testing.T) {
	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	parent := context.Background()

	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Go-routines", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)

	fmt.Println("Total GO Routine", runtime.NumGoroutine())

}
