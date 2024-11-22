package main

import (
	"fmt"
	"os"
)

func test(strs ...string) {
	fmt.Println(strs)
}

func main() {
	slice := os.Args[1:]
	fmt.Println(slice[2])
	fmt.Printf("%T\n", slice[2])
	slice = append(slice, "inha")
	fmt.Println(slice)
	test("123", "456")
}
