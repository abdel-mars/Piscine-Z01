package piscine

func IsUpper(s string) bool {
	res := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if res[i] < 65 || res[i] > 90 {
			return false
		}
	}
	return true
}
