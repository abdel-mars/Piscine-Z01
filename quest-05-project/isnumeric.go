package piscine

func IsNumeric(s string) bool {
	res := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if res[i] < '0' || res[i] > '9' {
			return false
		}
	}
	return true
}
