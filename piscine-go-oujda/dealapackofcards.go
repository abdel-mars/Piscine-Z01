package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// deck = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < 12; i += 3 {
		player := i/3 + 1
		end := i + 3
		fmt.Printf("Player %d: ", player)
		for i, card := range deck[i:end] {
			fmt.Printf("%d", card)
			if i != 2 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
