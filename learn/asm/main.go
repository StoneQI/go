package main

import "fmt"

type ints []int
func (i *ints) Sum() int {
	s := 0
	for _, v := range {
	  s += v
	}
	return s
}

func main() {

	type Summer interface {
		Sum() int
	}

	

	



	var i ints
	var s Sumer = i // nil value can satisfy interface
	fmt.Println(s==nil, s.Sum()) // true, 0

}
