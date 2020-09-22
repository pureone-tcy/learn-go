package main

import (
	"fmt"
	"time"
)

// Default Selection
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			// return
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("###### END ######")
}
