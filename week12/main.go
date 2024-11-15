package main

import (
	"fmt"
	"time"
)

func main() {
	var dates = [3]time.Time{time.Unix(1, 0),
		time.Unix(12, 0),
		time.Unix(1735244111, 0)}

	for i, _ := range dates {
		fmt.Println(i)
	}
	//for i := 0; i < len(dates); i++ {
	//	fmt.Println(i, "번째 배열", dates[i])
	//}
}
