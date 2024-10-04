package main

import (
	"fmt"
)

func main() {

	i := 13
	f := 12.9

	fmt.Printf("Value i : %d\n value f : %f", i, f)
	fmt.Printf("%d * %f = %f", i, f, float64(i)*f)
}
