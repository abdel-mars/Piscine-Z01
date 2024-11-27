package piscine

func BasicAtoi(s string) int {
	res := 0
	for i := 0; len(s) >= 0; i++ {
		res = res*10 + (int(s[i]) - 48)
	}
	return res
}
