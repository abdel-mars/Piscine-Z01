package piscine

func add(str1, sep, str2 string) string {
	if str1 == "" {
		return str2
	}
	return str1 + sep + str2
}

func Join(strs []string, sep string) string {
	res := ""
	for i := 0; len(strs) > i; i++ {
		res = add(res, sep, strs[i])
	}
	return res
}
