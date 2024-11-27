package piscine

func BasicAtoi2(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = res*10 + (int(s[i]) - 48)
	}
	return res
}
