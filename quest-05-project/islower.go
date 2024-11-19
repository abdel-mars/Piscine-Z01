package piscine

func IsLower(s string) bool {
	res := []rune(s)
	for i := 0; i < len(s); i++ {
		if res[i] < 97 || res[i] > 122 {
			return false
		}
	}
	return true
}
