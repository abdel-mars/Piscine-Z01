package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := 0; len(args) > i; i++ {
		if args[i] == "galaxy" || args[i] == "01" || args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
