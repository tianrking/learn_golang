package main

import "fmt"

func main() {
	GG := [10]string{""}
	for index, value := range GG {
		fmt.Println(index, value)
	}
}
