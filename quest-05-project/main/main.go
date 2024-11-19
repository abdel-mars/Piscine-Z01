package main

import (
	// "github.com/01-edu/z01"
	"piscine"
	"fmt"

)


func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, "HOLE"))
}
