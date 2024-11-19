package piscine

func IsAlpha(s string) bool {
	res := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if res[i] < 'a' || res[i] > 'z' {
			if res[i] < 'A' || res[i] > 'Z' {
				if res[i] < '0' || res[i] > '9' {
					return false
				}
			}
		}
	}
	return true
}
