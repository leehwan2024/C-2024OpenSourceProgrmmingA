package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// isPrime 함수: 주어진 수가 소수인지 판별
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
				return false
			}
		}
	}
	return true
}

// getInteger 함수: 사용자로부터 정수를 입력받음
func getInteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	// 첫 번째 점수 입력 받기
	fmt.Print("첫 번째 점수 입력: ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	// 두 번째 점수 입력 받기
	fmt.Print("두 번째 점수 입력: ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	// 첫 번째 점수부터 두 번째 점수까지 소수 출력
	fmt.Printf("소수 리스트 (%d부터 %d까지):\n", n1, n2)
	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // 마지막 줄바꿈
}
