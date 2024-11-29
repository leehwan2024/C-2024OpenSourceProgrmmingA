package main

import "fmt"

type student struct {
	id   int64
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444010
	student1.name = "honggildong"
	student1.gpa = 3.9
	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 2024440030
	student2.name = "choijihu"
	student2.gpa = 3.5
	fmt.Println(student2.name, ":", student2.gpa)
}
