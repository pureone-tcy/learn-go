package main

import (
	"fmt"
	"regexp"
)

// Regex
func main() {
	// Simple
	match, _ := regexp.MatchString("a([a-z0-9]+)e", "apple")
	fmt.Println(match)

	// Optimization
	r := regexp.MustCompile("a([a-z0-9]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// Return Match string
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	// Return Match string slice
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss)
}
