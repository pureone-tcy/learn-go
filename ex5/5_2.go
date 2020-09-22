package main

import (
	"fmt"
	"time"
)

// channel
// - blocking
// - queueing
// main <- ch1 -> goroutine1
// main <- ch1 -> goroutine2
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(5000 * time.Millisecond)
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 1, 1, 1, 1}
	c := make(chan int)
	go goroutine1(s1, c) // Parallel func running
	go goroutine2(s2, c) // Parallel func running
	x := <-c             // Blocking c <- sum
	fmt.Println(x)       // Output
	y := <-c             // Blocking c <- sum
	fmt.Println(y)       // Output
}
