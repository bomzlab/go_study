package main

import (
	"fmt"
	"time"
)

func main13() {
	tick := time.Tick(5 * time.Second)
	boom := time.After(5 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("     default")
			time.Sleep(5 * time.Millisecond)

		}
	}
}
