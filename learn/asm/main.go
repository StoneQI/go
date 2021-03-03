package main

import "fmt"

func main() {

	var w []int
	w1 := make([]int, 0)

	if w == nil {
		fmt.Println("w is nil")
	}
	if w1 == nil {
		fmt.Println("w is nil")
	}

}
0x002f 00047 (/workspaces/go/learn/asm/main.go:7)       MOVQ    $0, "".w+112(SP)
0x0038 00056 (/workspaces/go/learn/asm/main.go:7)       XORPS   X0, X0
0x003b 00059 (/workspaces/go/learn/asm/main.go:7)       MOVUPS  X0, "".w+120(SP)
0x0040 00064 (/workspaces/go/learn/asm/main.go:8)       LEAQ    ""..autotmp_3+48(SP), AX
0x0045 00069 (/workspaces/go/learn/asm/main.go:8)       TESTB   AL, (AX)
0x0047 00071 (/workspaces/go/learn/asm/main.go:8)       JMP     73
0x0049 00073 (/workspaces/go/learn/asm/main.go:8)       MOVQ    AX, "".w1+88(SP)
0x004e 00078 (/workspaces/go/learn/asm/main.go:8)       XORPS   X0, X0
0x0051 00081 (/workspaces/go/learn/asm/main.go:8)       MOVUPS  X0, "".w1+96(SP)
