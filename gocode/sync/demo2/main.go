package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x  int
	wg sync.WaitGroup
	//mux    sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock()
	x = x + 1
	rwlock.Unlock()
	time.Sleep(time.Millisecond * 10)
	wg.Done()

}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()

	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
