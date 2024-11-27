package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	var items string
	for _, n := range str {
		if n == ' ' {
			list[items] += 1
			items = ""
		} else if n != ' ' {
			items += string(byte(n))
		}
	}
	list[items] += 1
	return list
}
