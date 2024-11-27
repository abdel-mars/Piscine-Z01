package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintS(s string) {
	for i := 0; len(s) > i; i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	args := os.Args

	for i := len(args) - 1; i > 0; i-- {
		PrintS(args[i])
		z01.PrintRune('\n')
	}
}
