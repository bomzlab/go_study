package main

import (
	"fmt"
	"time"
)

func main8() {
	fmt.Println("main call")

	done := make(chan bool)

	go longTime2(done)
	go shortTime2(done)

	fmt.Println("end :: ", <-done)
	fmt.Println("end :: ", <-done)

	fmt.Println("main end>>")
}

func longTime2(done chan bool) {
	fmt.Println("long start : ", time.Now())

	time.Sleep(3 * time.Second)

	fmt.Println("long end : ", time.Now())

	done <- true
}

func shortTime2(done chan bool) {
	fmt.Println("short start : ", time.Now())

	time.Sleep(1 * time.Second)

	fmt.Println("short end : ", time.Now())

	done <- false
}
