package main

import "fmt"

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	gpa_slice[2] = 2.71
	gpa_slice = append(gpa_slice, 4.3)
	fmt.Println(gpa_slice, gpa)
}
