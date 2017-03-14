package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var socket net.Conn

func main19() {
	index := 0

	for {
		if !connect("127.0.0.1", 8080) {
			continue
		}

		write(&index)

		time.Sleep(3 * time.Second)
	}

}

func write(i *int) {
	msg := "good" + strconv.Itoa(*i)

	if _, err := socket.Write([]byte(msg)); err != nil {
		fmt.Println("전송 실패로 연결 종료")
		disconnect()
	} else {
		*i++
	}

}

func disconnect() {
	if socket != nil {
		socket.Close()
		socket = nil
	}
}

func connect(ip string, port int) bool {
	if socket != nil {
		return true
	}

	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("client connect fail : ", ip, ":", port)
		time.Sleep(1 * time.Second)
		return false
	}

	socket = conn
	return true
}
