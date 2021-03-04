package main

import "fmt"

func main() {

	type Summer interface {
		Sum() int
	}

	

	

	type ints []int
	func (i *ints) Sum() int {
		s := 0
		for _, v := range i{
		  s += v
		}

		return s
	  }

	var i ints
	var s Sumer = i // nil value can satisfy interface
	fmt.Println(s==nil, s.Sum()) // true, 0

}
