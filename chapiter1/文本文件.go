package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/k8s/learn_golang/chapiter1/文本文件.go")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
