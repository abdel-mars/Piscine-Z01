package piscine

func ConcatSep(str1 string, sep string, str2 string) string {
	if str1 == "" {
		return str2
	}
	return str1 + sep + str2
}

func Join(strs []string, sep string) string {
	var res string
	for i := 0; i < len(strs); i++ {
		res = ConcatSep(res, sep, strs[i])
	}
	return res
}
