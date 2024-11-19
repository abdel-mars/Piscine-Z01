package piscine

func ToLower(s string) string {
	res := []rune(s)
	for i := 0; i < len(s); i++ {
		if IsUpper(string(s[i])) {
			res[i] = rune(res[i] + 32)
		}
	}
	return string(res)
}
