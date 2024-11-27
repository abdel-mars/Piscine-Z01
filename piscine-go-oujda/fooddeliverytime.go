package piscine

type food struct {
	order    string
	preptime int
}

func SetTime(ptr *food) {
	if ptr.order == "burger" {
		ptr.preptime = 15
	} else if ptr.order == "chips" {
		ptr.preptime = 10
	} else if ptr.order == "nuggets" {
		ptr.preptime = 12
	} else {
		ptr.preptime = 404
	}
}

func FoodDeliveryTime(order string) int {
	food := &food{}
	food.order = order
	SetTime(food)
	return food.preptime
}
