package main

import "fmt"

func main() {
	var result string
	result = fmt.Sprintf("%0.2f 나누기 %0.2f는  %0.2f\n", 1.0, 3.0, 1.0/3.0)
	fmt.Println(result)
}
