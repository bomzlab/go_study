package main

import (
	"fmt"
	"sync"
	"time"
)

var list = []int{}

var lock = sync.Mutex{}

func main17() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			running(i)
		}(i)
	}

	wg.Wait()

	for i, v := range list {
		fmt.Println(i, " ranking is ", v)
	}

}

func running(i int) {
	for x := 1; x <= 2; x++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, " >> ", x, " running ...")
	}

	lock.Lock()
	list = append(list, i)
	lock.Unlock()
}
