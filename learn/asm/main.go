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

0x0025 00037 (/workspaces/go/learn/asm/main.go:12)      MOVQ    $0, "".a+24(SP)
0x002e 00046 (/workspaces/go/learn/asm/main.go:14)      MOVQ    $0, ""..autotmp_2+32(SP)
0x0037 00055 (/workspaces/go/learn/asm/main.go:14)      LEAQ    go.itab.*strings.Reader,io.Reader(SB), AX
0x003e 00062 (/workspaces/go/learn/asm/main.go:14)      MOVQ    AX, "".w+40(SP)
0x0043 00067 (/workspaces/go/learn/asm/main.go:14)      MOVQ    $0, "".w+48(SP)