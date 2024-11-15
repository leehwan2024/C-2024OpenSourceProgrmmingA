package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {

	var gpa [3]float64
	for i := 0; i < len(gpa); i++ {
		fmt.Println("Input float number:")
		gpa[i], _ = keyboard.GetFloat()
	}
	for i, v := range gpa {
		fmt.Printf("%v: %f\n", i, v)
	}
	/*var dates = [3]time.Time{time.Unix(1, 0),
		time.Unix(12, 0),
		time.Unix(1735244111, 0)}

	for i, _ := range dates {
		fmt.Println(i)
	}*/
	//for i := 0; i < len(dates); i++ {
	//	fmt.Println(i, "번째 배열", dates[i])
	//}
}
