package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printS(s string) {
	for i := 0; len(s) > i; i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	args := os.Args

	for i := 1; len(args) > i; i++ {
		printS(args[i])
		z01.PrintRune('\n')
	}
}
