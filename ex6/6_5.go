package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ioutil
func main() {
	content, err := ioutil.ReadFile("ex6/6_5.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	//if err := ioutil.WriteFile("sample.go", content, 0666); err != nil {
	//	return
	//}

}
