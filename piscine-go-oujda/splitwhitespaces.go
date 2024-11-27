package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	var land string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			land += string(s[i])
		} else if len(land) > 0 {
			res = append(res, land)
			land = ""
		}
	}
	if len(land) > 0 {
		res = append(res, land)
	}

	return res
}
