package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init call")
}

func main7() {
	fmt.Println("main call")

	go longTime()
	go shortTime()

	time.Sleep(5 * time.Second)

	fmt.Println("main end>>", time.Second)
}

func longTime() {
	fmt.Println("long start : ", time.Now())

	time.Sleep(3 * time.Second)

	fmt.Println("long end : ", time.Now())
}

func shortTime() {
	fmt.Println("short start : ", time.Now())

	time.Sleep(1 * time.Second)

	fmt.Println("short end : ", time.Now())
}
