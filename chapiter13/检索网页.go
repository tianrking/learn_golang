package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// response, err := http.Get("http://baidu.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(response)

	// // defer让函数或语句可以在当前函数执行完毕后
	// //（包括通过return正常结束或者panic导致的异常结束）执行。
	// // defer 语句通常用于一些成对操作的场景：打开连接/关闭连接
	// defer response.Body.Close()
	// // defer fmt.Println("GGG")
	// // fmt.Println("11")
	// // go fmt.Println("22")
	// // fmt.Println("33")
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))
	go responseSize("https://baidu.com/")
	time.Sleep(1 * time.Second) // 等待完成网络请求

}

func responseSize(url string) {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}
