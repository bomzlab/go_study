package main

import (
	"fmt"
	"net"
	"time"
)

var listener net.Listener

func main18() {

	for {
		if !openServer(":8080") {
			continue
		}

		fmt.Println("waiting client ...")
		client, err := listener.Accept()

		if err != nil {
			closeServer()
			continue
		}

		fmt.Println("client :: ", client)

		go handleClient(client)
	}

}

func closeServer() {
	listener.Close()
}

func openServer(port string) bool {
	if listener != nil {
		return true
	}

	ln, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("server open error :: ", err)
		time.Sleep(3 * time.Second)
		return false
	}

	listener = ln
	fmt.Printf("server open success. port %s\n", port)

	return true
}
