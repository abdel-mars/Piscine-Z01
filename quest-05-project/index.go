package piscine

func Index(s string, toFind string) int {
	for i := 0; len(s)-len(toFind) >= i; i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
