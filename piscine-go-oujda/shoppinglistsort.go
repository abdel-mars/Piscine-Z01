package piscine

func ShoppingListSort(slice []string) []string {
	var res []string
	for i := 0; len(slice) > i-1; i++ {
		for j := 0; len(slice)-i-1 > j; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
		res = slice
	}
	return res
}
