package main

import "fmt"

// channel
func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	c <- 1
	close(c)
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s1))
	go goroutine3(s1, c)

	for cc := range c {
		fmt.Println(cc)
	}
}
