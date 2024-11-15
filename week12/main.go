package main

import (
	"fmt"
	"time"
)

func main() {
	var dates = [3]time.Time{time.Unix(1, 0),
		time.Unix(12, 0),
		time.Unix(1976299143, 0)}
	fmt.Println(dates[0], dates[1], dates[2])

}
