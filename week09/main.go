package main

import "fmt"

func combie(i float64, j float64) float64 {
	a := i * j
	return a
}
func main() {
	var a, b float64
	a = combie(6.5, 8.6)
	b = combie(3.5, 1.6)
	fmt.Printf("%f\n %f\n", a, b)
}
