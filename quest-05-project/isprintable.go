package piscine

func IsPrintable(s string) bool {
	res := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if res[i] < 33 || res[i] > 126 {
			if res[i] < 130 || res[i] > 140 {
				if res[i] < 158 || res[i] > 255 {
					return false
				}
			}
		}
	}
	return true
}
