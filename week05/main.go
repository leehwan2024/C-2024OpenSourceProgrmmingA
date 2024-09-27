package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	var f float32 = 4.3 //f:=4.3

	fmt.Println(reflect.TypeOf(f))
	i := 55
	fmt.Println(reflect.TypeOf(i))

	fmt.Printf("%s\n", strings.Title("inhatc"))
	fmt.Println(math.Ceil(3.99))

	fmt.Printf("Value i : %d\n", i)
	fmt.Println("Value i : ", i)
	fmt.Println("Value i : ", i)
}
