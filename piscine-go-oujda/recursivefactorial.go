package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb >= 1 && nb <= 20 {
		return nb * IterativeFactorial(nb-1)
	}
	return 0
}
