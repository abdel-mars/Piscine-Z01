package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	i := 0
	n := len(s)

	for i < n {

		for s[i] == ' ' {
			i++
		}
		if s[i] == '-' || s[i] == '+' {
			if s[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && (s[i] >= '0' && s[i] <= '9') {
			res = res*10 + int(s[i]-48)
			i++
		}
		i++
	}
	return res * sign
}
