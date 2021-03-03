package main

import "fmt"

func main() {

	var w chan int       // nil map
	w1 := make(chan int) // empty map

	if w == nil {
		fmt.Println("w is nil")
	}
	if w1 == nil {
		fmt.Println("w is nil")
	}

}
