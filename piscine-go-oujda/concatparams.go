package piscine

func ConcatSep(str1 string, sep string, str2 string) string {
	if str1 == "" {
		return str2
	}
	return str1 + sep + str2
}

func ConcatParams(args []string) string {
	var res string
	for i := 0; i < len(args); i++ {
		res = ConcatSep(res, "\n", args[i])
	}
	return res
}
