package piscine

func StringToIntSlice(str string) []int {
	var res []int
	runes := []rune(str)
	for i := 0; len(runes) > i; i++ {
		res = append(res, int(runes[i]))
		// res += int(runes[i])
	}
	return res
}
