package main

import "fmt"

func main() {

	var w map[int]string       // nil map
	w1 := make(map[int]string) // empty map

	_, _ = w[1]
	if w == nil {
		fmt.Println("w is nil")
	}
	if w1 == nil {
		fmt.Println("w is nil")
	}

}
