package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now.Unix())
	target := rand.Intn(2) + 1
	fmt.Printf("%d\n", target)
}
