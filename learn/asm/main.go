package main

import (
	"bytes"
	"io"
)

func for1() int {
	a := 0
	for i := 0; i < 5; i = i + 1 {
		a = a + i
	}
	return a
}

func main() {
	var aa = func() {
		a := 0
		for i := 0; i < 5; i = i + 1 {
			a = a + i
		}
	}

	_ = aa
	var w io.Writer

	w = bytes.Buffer()
	_ = w
	var bb map[string]int
	_ = bb
	_ = for1()
}
