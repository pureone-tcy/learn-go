package main

import (
	"context"
	"fmt"
	"time"
)

// context
// - cancel
func longProcess(_ context.Context, ch chan<- string) {
	fmt.Println("run")
	time.Sleep(4 * time.Second)
	fmt.Println("finish")

	ch <- "Result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//ctx := context.TODO()

	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("Success")
			break CTXLOOP
		}
	}
	fmt.Println("### End ###")
}
