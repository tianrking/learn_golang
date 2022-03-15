package main

import (
	"fmt"
	"time"
)

type num_struct struct {
	old_num int
	new_num int
}

func main() {
	fmt.Println("Start")
	test_array := []int{10, 20, 30, 40}
	my_channel := make(chan num_struct)
	for _, index := range test_array {
		go test_gg(index, my_channel)
	}
	time.Sleep(3 * time.Second)
	for i := len(test_array); i > 0; i-- {
		fmt.Println(<-my_channel)
	}
}

func test_gg(_num int, channel chan num_struct) {
	time.Sleep(2 * time.Second)
	fmt.Println(_num)
	channel <- num_struct{old_num: _num, new_num: _num + 1}
}
