package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("점수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	//count := 0
	var isPrime bool = true

	if n <= 1 {
		isPrime = false
	}

	j := 0
	for j = 2; j <= int(math.Sqrt(float64(n))); j++ {
		if n%j == 0 {
			//count++
			isPrime = false
			break
		}
		fmt.Printf("%d번 반복했습니다\n", j-1)
	}
	if isPrime {
		fmt.Printf("%d은(는) 소수입니다.", n)
	} else {
		fmt.Printf("%d은(는) 소수가 아닙니다.", n)
	}
}
