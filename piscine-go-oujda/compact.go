package piscine

func Compact(ptr *[]string) int {
	res := []string(nil)

	for _, str := range *ptr {
		if str != "" {
			res = append(res, str)
		}
	}
	*ptr = res
	return len(res)
}
