package main

import (
	"io"
	"strings"
)

func main() {

	var w io.Reader

	w = strings.NewReader("1234")

	_ = w

}
