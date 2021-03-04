package main

import "fmt"

type ints []int

type Summer interface {
	Sum() int
}

func (i ints) Sum() int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func main() {

	var b = make(map[string]string)
	fmt.Println(b["ccc"])
	fmt.Println(b["ddd"])
	var i ints
	var s Summer = i
	var c Summer
	fmt.Println(c)
	fmt.Println(s)                 // nil value can satisfy interface
	fmt.Println(s == nil, s.Sum()) // true, 0

}
