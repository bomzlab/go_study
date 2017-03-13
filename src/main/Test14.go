package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const MAX_LOOP = 1000

type counter struct {
	i    int64
	lock sync.Mutex
	once sync.Once
}

func (c *counter) increment() {

	// 1번만 실행하는 락
	c.once.Do(func() {
		fmt.Println("once call")
	})

	atomic.AddInt64(&c.i, 1) // 해당 처리가 될때는 시간당 cpu 분할 처리를 하지 않는다

	//	sync.Mutex 를 이용한 방법
	//	c.lock.Lock()
	//	c.i += 1
	//	c.lock.Unlock()
}

func (c *counter) display() {
	fmt.Println("display :: ", c.i)
}

func main14() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("my cpu :: ", runtime.NumCPU())

	c := counter{i: 0}

	// 비효율적인 방법
	//	ch := make(chan struct{})
	//
	//	for i := 0; i < MAX_LOOP; i++ {
	//		go func() {
	//			c.increment()
	//			ch <- struct{}{}
	//		}()
	//	}
	//
	//	for i := 0; i < MAX_LOOP; i++ {
	//		<-ch
	//	}

	// 효율적인 방법
	wg := sync.WaitGroup{}

	for i := 0; i < MAX_LOOP; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			c.increment()
		}()
	}

	wg.Wait()

	c.display()

}
