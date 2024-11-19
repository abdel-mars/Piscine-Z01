package piscine

func NRune(s string, n int) rune {
	var empty rune
	if n > len(s) {
		return empty
	}
	for i := 0; i <= n-1; i++ {
		if i == n-1 {
			return []rune(s)[i]
		}
	}
	return empty
}
