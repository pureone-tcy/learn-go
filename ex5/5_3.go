package main

import "fmt"

// buffered channel
// - buffered
// - range close
func main() {
	ch1 := make(chan int, 2)
	ch1 <- 100
	fmt.Println(len(ch1))
	ch1 <- 200
	fmt.Println(len(ch1))
	_, _ = <-ch1, <-ch1
	fmt.Println(len(ch1))

	ch2 := make(chan int, 2)
	ch2 <- 100
	ch2 <- 200
	close(ch2)
	for c := range ch2 {
		fmt.Println(c)
	}
}
