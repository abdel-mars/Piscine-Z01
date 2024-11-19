package piscine

func LastRune(s string) rune {
	var res rune
	for i := 0; len(s)-1 >= i; i++ {
		if i == len(s)-1 {
			return []rune(s)[i]
		}
	}
	return res
}
