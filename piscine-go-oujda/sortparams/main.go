package main

import (
	"os"

	"github.com/01-edu/z01"
)

//	func Swap(a int, b int) {
//		a, b = b, a
//	}
func main() {
	args := os.Args
	var count int
	for str := range args {
		count = str + 1
	}
	for c1 := 1; c1 < count; c1++ {
		for c2 := c1 + 1; c2 < count; c2++ {
			if args[c1] > args[c2] {
				args[c1], args[c2] = args[c2], args[c1]
			}
		}
	}
	for c2 := 1; c2 <= count-1; c2++ {
		for _, str := range args[c2] {
			z01.PrintRune(str)
		}
		z01.PrintRune('\n')
	}
}
