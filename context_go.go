package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("### Go Context ###")

	ctx := context.Background()
	fmt.Printf("background\t%v\n", ctx)
	fmt.Printf("errors\t\t%v\n", ctx.Err())
	fmt.Printf("type\t\t%T\n", ctx)
	fmt.Printf("----------------------------------\n")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Printf("background\t%v\n", ctx)
	fmt.Printf("errors\t\t%v\n", ctx.Err())
	fmt.Printf("errors\t\t%T\n", ctx.Err())
	fmt.Printf("type\t\t%T\n", ctx)
	fmt.Printf("cancel\t\t%v\n", cancel)
	fmt.Printf("type\t\t%T\n", cancel)

	fmt.Printf("cancel----------------------------\n")
	cancel()
	fmt.Printf("background\t%v\n", ctx)
	fmt.Printf("errors\t\t%v\n", ctx.Err())
	fmt.Printf("errors\t\t%T\n", ctx.Err())
	fmt.Printf("type\t\t%T\n", ctx)
	fmt.Printf("cancel\t\t%v\n", cancel)
	fmt.Printf("type\t\t%T\n", cancel)

	fmt.Printf("----------------------------------\n")

	contextExample()
}

func contextExample() {
	// would be main.go
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("error check 1   :", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return  // return not to leak a goroutine
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2   :", ctx.Err())
	fmt.Println("num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 3   :", ctx.Err())
	fmt.Println("num goroutines 3:", runtime.NumGoroutine())
}
