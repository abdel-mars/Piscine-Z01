package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	tmp := 0
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			tmp++
		}
	}
	if tmp == 0 {
		return true
	}
	return false
}
