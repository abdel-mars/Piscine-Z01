package piscine

func BasicJoin(elems []string) string {
	res := ""
	for i := 0; i < len(elems); i++ {
		res = Concat(res, elems[i])
	}
	return res
}
