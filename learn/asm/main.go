package main

import "fmt"

func main() {

	var w chan int       // nil map
	w1 := make(chan int) // empty map

	if w == nil {
		fmt.Println("w is nil")
	}
	if w1 == nil {
		fmt.Println("w is nil")
	}

}

MOVQ    $0, "".w+56(SP)
        0x0038 00056 (/workspaces/go/learn/asm/main.go:8)       LEAQ    type.chan int(SB), AX
        0x003f 00063 (/workspaces/go/learn/asm/main.go:8)       MOVQ    AX, (SP)
        0x0043 00067 (/workspaces/go/learn/asm/main.go:8)       MOVQ    $0, 8(SP)
        0x004c 00076 (/workspaces/go/learn/asm/main.go:8)       PCDATA  $1, $1
        0x004c 00076 (/workspaces/go/learn/asm/main.go:8)       CALL    runtime.makechan(SB)
        0x0051 00081 (/workspaces/go/learn/asm/main.go:8)       MOVQ    16(SP), AX
        0x0056 00086 (/workspaces/go/learn/asm/main.go:8)       MOVQ    AX, "".w1+48(SP)
