package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calc_cost(visitors []visitor) int {
	totalCost := 0
	for _, v := range visitors {
		totalCost += v.cost
	}
	return totalCost

}

func main() {
	var numvisitors int
	fmt.Print("How many visitors?: ")
	fmt.Scanln(&numvisitors)

	var vs []visitor
	vs = make([]visitor, numvisitors)

	for i := 0; i < numvisitors; i++ {
		var age int
		fmt.Printf("input age: ")
		fmt.Scanln(&age)
		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default: // if over age 65
			vs[i] = visitor{age: age, cost: 7000}
		}
	}
	fmt.Printf("Total cost is %d won.", calc_cost(vs))
}
