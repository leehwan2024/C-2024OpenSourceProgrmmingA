package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Input your name")
	name, err := in.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)
	fmt.Println(err)

}
