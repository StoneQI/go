package main

import "fmt"

func main() {

	var w map[int]string          // nil map
	w1 := make(map[int]string, 0) // empty map

	if w == nil {
		fmt.Println("w is nil")
	}
	if w1 == nil {
		fmt.Println("w is nil")
	}

}
