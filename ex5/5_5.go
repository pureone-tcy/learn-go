package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer, Consumer
func producer(ch chan int, i int) {
	// something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			// do something
			fmt.Printf("process:[%d]\n", i)
		}()
	}
	fmt.Println("##################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
