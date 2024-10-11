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
	score, err := strconv.ParseInt(i, 10, 32) // 문자열 변수 i의 값을 정수형(32비트)로 변환, 입력받은 값-> 16진수로변환
	//                         실수형 받고 싶으면 ParseFloat(i,32) ㄱㄱ
	var aOrnot string
	if score >= 95 {
		aOrnot = "A"

	} else {
		aOrnot = "BCDF"

	}
	fmt.Printf("%d점은 %s등급입니다.", score, aOrnot)
}
