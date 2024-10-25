package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	answer := rand.Intn(100) + 1
	fmt.Printf("%d\n", answer)

	fmt.Print("점수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	guess, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답!")
	} else if answer > guess {
		fmt.Println("입력 수가 정답보다 작습니다")
	} else if answer < guess {
		fmt.Println("입력 수가 정답보다 큽니다")
	}
	fmt.Printf("정답은 %d", answer)

}
