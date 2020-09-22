package main

import "fmt"

// iota
const (
	c1 = iota
	c2
	c3
)

// ex6. KB,MB
const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
