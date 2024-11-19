package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) > 2 {
		printStr("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		printStr("File name missing")
		return
	} else {
		readfile, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			printStr("Wrong File Name")
		}
		fmt.Printf("%s", readfile)
	}
}
