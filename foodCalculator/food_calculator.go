package main

import (
	"fmt"
)

type FoodCalculator struct {
	Order  []OrderDetail
	Member bool
}

type OrderDetail struct {
	Item    string
	Quntity int
}

var foodList = map[string]int{
	"Red":    50,
	"Green":  40,
	"Blue":   30,
	"Yellow": 50,
	"Pink":   80,
	"Purple": 90,
	"Orange": 120,
}

var discountFood = []string{"Orange", "Pink", "Green"}

func (fc *FoodCalculator) Calculator() float32 {
	var total float32

	for _, f := range fc.Order {
		if price, ok := foodList[f.Item]; ok {

			// getting discount on Orange, Pink or Green
			if contains(discountFood, f.Item) {
				// get a 5% discount for each bundles
				before := price * f.Quntity
				bundles := f.Quntity / 2
				discount := 2 * float32(price) * float32(bundles) * 0.05
				total += float32(before) - discount

			} else {
				total += float32(price) * float32(f.Quntity)
			}

		} else {
			fmt.Println("food not found: ", f.Item)
			continue
		}

	}

	if fc.Member {
		total = total * 0.9
	}

	return total

}

func contains(arr []string, x string) bool {
	for _, n := range arr {
		if n == x {
			return true
		}
	}
	return false
}

func main() {
	case1 := FoodCalculator{Order: []OrderDetail{{"Purple", 8}, {"Red", 1}}, Member: true}
	fmt.Println(case1.Calculator())
}
