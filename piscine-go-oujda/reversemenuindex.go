package piscine

func ReverseMenuIndex(menu []string) []string {
	res := make([]string, len(menu))
	for i, j := 0, len(menu)-1; i < len(res); i, j = i+1, j-1 {
		res[i] = menu[j]
	}
	return res
}
