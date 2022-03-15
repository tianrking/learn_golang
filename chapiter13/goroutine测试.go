package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	test_array := []int{10, 20, 30, 40}

	my_channel := make(chan int)
	for _, index := range test_array {

		go test_gg(index, my_channel)
	}

	time.Sleep(3 * time.Second)

	fmt.Println(<-my_channel)
	fmt.Println(<-my_channel)
	fmt.Println(<-my_channel)
	fmt.Println(<-my_channel)
}

func test_gg(_num int, channel chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println(_num)

	channel <- _num + 1
}
