package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("Accept failed: %v", err)
			continue
		}
		go process(client)
	}
}

func process(client net.Conn) {
	remoteAddr := client.RemoteAddr().String()
	fmt.Printf("Connection from %s\n", remoteAddr)
	client.Write([]byte("Hello world!\n"))
	client.Close()
}
