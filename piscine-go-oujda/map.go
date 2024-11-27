package piscine

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, len(a))
	for i, s := range a {
		res[i] = f(s)
	}
	return res
}
