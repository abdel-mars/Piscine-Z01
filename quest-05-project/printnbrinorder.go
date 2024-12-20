package piscine

import "github.com/01-edu/z01"

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func SortIntegerTable(table []int) {
	for i := range table {
		for j := range table {
			if table[i] < table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
	}
	numbers := []int{}
	for n > 0 {
		mod := n % 10
		numbers = append(numbers, mod)
		n /= 10
	}
	SortIntegerTable(numbers)
	for i := 0; i < len(numbers); i++ {
		z01.PrintRune(rune(numbers[i] + 48))
	}
}
