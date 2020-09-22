package main

import (
	"fmt"
	"sync"
)

// goroutine
// - basic
// - wait group
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	defer wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	// time.Sleep(200 * time.Millisecond)
	wg.Wait()
}
