package main

import (
	"fmt"
	"time"
)

func main() {
	var dates = [3]time.Time{time.Unix(1, 0),
		time.Unix(12, 0),
		time.Unix(1735244111, 0)}

	for i := 0; i < len(dates); i++ {
		fmt.Println(i, "번째 배열", dates[i])
	}
	// fmt.Printf("%v\n", dates)
	// fmt.Println(dates[0], dates[1], dates[2])

}
