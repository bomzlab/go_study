package main

import (
	"fmt"
)

func main12() {

	ch := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("ch::", <-ch)
		}

		quit <- 0
	}()

	check(ch, quit)
}

func check(ch, quit chan int) {
	x := 1

	for {
		select {
		case ch <- x:
			x++
		case <-quit:
			fmt.Println("quit :: ")
			return
		default:
			fmt.Println("default call")
		}
	}
}
