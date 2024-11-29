package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("Input your name (or exit to 'q'): ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("your age?")
		fmt.Scanln(&age)

		ages[name] = age

	}
	for name, age := range ages {
		fmt.Printf("%s : %d years old\n", name, age)
	}
}
