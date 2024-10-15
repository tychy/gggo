package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid os.Args length")
	}

	var argv []string
	argv = os.Args[1:]

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")
	a, _ := strconv.Atoi(argv[0])
	fmt.Printf("  mov rax, %d\n", a)
	fmt.Printf("  ret\n")
	return
}
