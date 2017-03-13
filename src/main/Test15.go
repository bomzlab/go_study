package main

import (
	"fmt"
	"runtime"
	"time"
)

func main15() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	quit := make(chan struct{})
	done := process15(quit)
	timeout := time.After(2 * time.Second)

	select {
	case d := <-done:
		fmt.Println("done >> ", d)
	case <-timeout:
		fmt.Println("timeout")
		quit <- struct{}{}
	}
}

func process15(quit chan struct{}) chan string {
	done := make(chan string)

	go func() {

		go func() {
			time.Sleep(10 * time.Second)
			done <- "Complate"
		}()

		<-quit
		fmt.Println("닫아라")
	}()

	return done
}
