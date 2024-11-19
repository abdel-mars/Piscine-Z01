package piscine

func AlphaCount(s string) int {
	res := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			res++
		}
		if s[i] <= 'A' && s[i] >= 'Z' || s[i] <= 'a' && s[i] >= 'z' {
			i++
		}
	}
	return res
}
