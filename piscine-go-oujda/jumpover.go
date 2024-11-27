package piscine

func JumpOver(str string) string {
	var res string
	newline := "\n"
	if len(str) == 0 {
		return newline
	}
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	res += newline
	return res
}
