package main

import (
	"fmt"
	"time"
)

// select
// - non Blocking
// - 2 channel, 2 goroutine
//   main <- channel1 -> goroutineA
//   main <- channel2 -> goroutineB
func goroutineA(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutineB(ch chan int) {
	for {
		ch <- 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan int)

	go goroutineA(c1)
	go goroutineB(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
