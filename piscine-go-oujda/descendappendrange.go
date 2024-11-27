package piscine

func DescendAppendRange(max, min int) []int {
	res := []int{}
	for i := max; i > min; i-- {
		res = append(res, i)
	}
	return res
}
