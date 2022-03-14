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
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)

	for gg := 10; gg > 0; gg-- {
		fmt.Println(gg)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		// num, err := strconv.ParseFloat(input, 64)
		num, err := strconv.ParseInt(input, 10, 10)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(num)
		guess, _ := strconv.Atoi(input)
		// fmt.Println(guess)

		if guess > target {
			fmt.Println("guess > target")
		} else if guess < target {
			fmt.Println("guess < target")
		} else {

			break
		}

		// fmt.Println(reflect.TypeOf(num)) int64
		// fmt.Println(reflect.TypeOf(target)) int
	}
}
