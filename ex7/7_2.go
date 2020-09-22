package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// semaphore
var s *semaphore.Weighted = semaphore.NewWeighted(2)

func longProcess(ctx context.Context) {

	// Other goroutine cancel
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}

	// goroutine thread block!!!
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	defer s.Release(1)

	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(4 * time.Second)
}
