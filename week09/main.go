package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)
	for guesses := 3; guesses > 0; guesses-- {
		fmt.Printf("%d번 기회가 남았습니다. \n점수 입력: ", guesses)
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
			break
		} else if (answer > guess || answer < guess) && guesses == 1 {
			fmt.Println("기회를 모두 소진했습니다. you lose")
		} else if answer > guess {
			fmt.Println("입력 수가 정답보다 작습니다")
		} else if answer < guess {
			fmt.Println("입력 수가 정답보다 큽니다")
		}

	}
}
