package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for j := 3; j*j <= n; j += 2 {
			if n%j == 0 {
				//count++
				return false
			}
			//fmt.Printf("%d\n", j)
		}
	}
	return false
}
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

	if isPrime(n) {
		fmt.Printf("%d은(는) 소수입니다.", n)
	} else {
		fmt.Printf("%d은(는) 소수가 아닙니다.", n)
	}
}
