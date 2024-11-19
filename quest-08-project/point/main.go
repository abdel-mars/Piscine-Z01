package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func setPoint(ptr *point) {
	ptr.x = 33
	ptr.y = 34
}

func PrintNbr(number int) {
	c := '0'
	u := '0'
	sign := 1

	if number < 0 {
		z01.PrintRune('-')
		sign = -1
	}
	res := number * sign
	for i := 0; i < res/10; i++ {
		c++
	}
	z01.PrintRune(c)
	for i := 0; i < res%10; i++ {
		u++
	}
	z01.PrintRune(u)
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	PrintNbr(points.x)
	printStr(", y = ")
	PrintNbr(points.y)
	printStr("\n")
}
