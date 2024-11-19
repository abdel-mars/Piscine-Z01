package piscine


func FirstRune(s string) rune {
	if len(s) == 0 {
		var empty rune
		return rune(empty)
	}
	return []rune(s)[0]
}
