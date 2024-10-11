package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 13
	var f float64 = 12.8
	c1 := '김'
	c2 := 'b'
	fmt.Printf("Value i : %d\n value f : %f", i, f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)

	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
