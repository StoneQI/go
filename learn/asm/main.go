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

0x001d 00029 (/workspaces/go/learn/asm/main.go:10)      XORPS   X0, X0
0x0020 00032 (/workspaces/go/learn/asm/main.go:10)      MOVUPS  X0, "".w+32(SP)
0x0025 00037 (/workspaces/go/learn/asm/main.go:12)      LEAQ    go.string."1234"(SB), AX
0x002c 00044 (/workspaces/go/learn/asm/main.go:12)      MOVQ    AX, (SP)
0x0030 00048 (/workspaces/go/learn/asm/main.go:12)      MOVQ    $4, 8(SP)
0x0039 00057 (/workspaces/go/learn/asm/main.go:12)      PCDATA  $1, $0
0x0039 00057 (/workspaces/go/learn/asm/main.go:12)      CALL    strings.NewReader(SB)
0x003e 00062 (/workspaces/go/learn/asm/main.go:12)      MOVQ    16(SP), AX
0x0043 00067 (/workspaces/go/learn/asm/main.go:12)      MOVQ    AX, ""..autotmp_1+24(SP)
0x0048 00072 (/workspaces/go/learn/asm/main.go:12)      LEAQ    go.itab.*strings.Reader,io.Reader(SB), CX
0x004f 00079 (/workspaces/go/learn/asm/main.go:12)      MOVQ    CX, "".w+32(SP)
0x0054 00084 (/workspaces/go/learn/asm/main.go:12)      MOVQ    AX, "".w+40(SP)