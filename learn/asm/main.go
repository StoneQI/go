package main

import (
	"io"
	"strings"
)

func main() {

	var w io.Reader

	var a *strings.Reader
	w = strings.NewReader("1234")

	_ = w

}

0x001d 00029 (/workspaces/go/learn/asm/main.go:10)      XORPS   X0, X0
0x0020 00032 (/workspaces/go/learn/asm/main.go:10)      MOVUPS  X0, "".w+32(SP)
0x0025 00037       
0x002c 00044       
0x0030 00048       
0x0039 00057       
0x0039 00057       
0x003e 00062       
0x0043 00067       
0x0048 00072       
0x004f 00079       
0x0054 00084       