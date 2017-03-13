package main

import (
	"fmt"
	"sync"
	"time"
)

func main11() {
	// error example

	ch := make(chan int)

	fmt.Println("start : ", ch)

	type lock struct {
		m sync.Mutex
	}

	lo := lock{}

	go func(lo *lock) {
		for i := 0; i < 10; i++ {
			lo.m.Lock()
			ch <- i
			lo.m.Unlock()
			time.Sleep(1 * time.Second)
		}
	}(&lo)

	for i := range ch {
		time.Sleep(5 * time.Second)

		lo.m.Lock()
		fmt.Println(">>>", <-ch, " <> ", i)
		lo.m.Unlock()

	}

	fmt.Println("end")

	close(ch)
}
