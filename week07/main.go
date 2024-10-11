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

	score, err := strconv.ParseInt(i, 16, 32) // 문자열 변수 i의 값을 정수형(32비트)로 변환, 입력받은 값-> 16진수로변환
	//                         실수형 받고 싶으면 ParseFloat(i,32) ㄱㄱ
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 {
		fmt.Println("A")
		fmt.Printf("%d", score)

	} else {
		fmt.Println("BCDE")
		fmt.Printf("%d", score)
	}
}
