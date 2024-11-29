package main

import "fmt"

func main() {
	var student1 struct {
		id   int64
		name string
		gpa  float32
	}
	student1.id = 202444010
	student1.name = "honggildong"
	student1.gpa = 3.9
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int64
		name string
		gpa  float32
	}
	student2.id = 2024440030
	student2.name = "choijihu"
	student2.gpa = 3.5
	fmt.Println(student2.name, ":", student2.gpa)
}
