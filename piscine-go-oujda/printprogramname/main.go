package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argname := []rune(os.Args[0])
	for i := range argname {
		if i > 1 {
			z01.PrintRune(argname[i])
		}
	}
	z01.PrintRune('\n')
}
