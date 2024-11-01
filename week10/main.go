package main

import (
	"bufio"
	"fmt"
	"log"
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

	count := 0
	for j := 1; j <= n; j++ {
		if n%j == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Printf("%d은(는) 소수입니다.", n)
	} else {
		fmt.Printf("%d은(는) 소수가 아닙니다.", n)
	}
}
