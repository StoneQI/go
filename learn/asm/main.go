package main

import (
	"io"
	"strings"
)



func main() {

	type Summer intface {
		Sum() int
	}
	  
	var t *int // nil of type *tree
	var s Summer = t // nil指针，可以是合法的interface类型的值
	  // 此时，对接接口类型变量s而言，其类型为*tree，值为nil，也就是说(*tree, nil)行的interface
	  
	fmt.Println(t==nil, s.Sum()) // true, 0
	  
	// type ints []int
	// func (i *ints) Sum() int {
	// 	s := 0
	// 	for _, v := range i{
	// 	  s += v
	// 	}

	// 	return s
	//   }
	  
	// var i ints
	// var s Sumer = i // nil value can satisfy interface
	// fmt.Println(s==nil, s.Sum()) // true, 0

}
