package main

import (
	"io"
	"strings"
)

func main() {

	var w io.Reader

	var a *strings.Reader

	w = a

	w = strings.NewReader("1234")

	_ = w

}
