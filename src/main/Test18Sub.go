package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleClient(client net.Conn) {

	defer func() {
		client.Close()
		fmt.Println("client call close")
	}()

	reader := bufio.NewReader(client)

	var data []byte = make([]byte, 300)

	for {
		//		size, err := client.Read(data)
		size, err := reader.Read(data)
		fmt.Println("client read size :: ", size, " <<>> ", len(data))

		if err != nil {
			fmt.Println("client close. ", err)
			break
		}

		if size < 0 {
			fmt.Println("client read size error : ", size)
			break
		}

		if size == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		fmt.Println("read size :: ", size, " , data :: ", string(data[:size]))
	}
}
