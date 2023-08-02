package main

// sync.Map 线程安全的map
import (
	"fmt"
	"sync"
)

var s map[int]int
var s2 = sync.Map{} //线程安全的map
var wg sync.WaitGroup

func getKey(i int) int {
	return s[i]
}

func setValue(i, j int) {
	s2.Store(i, j)
}

/*
	func main() {
		s = make(map[int]int)
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(i int) {
				setValue(i, i+100)
				fmt.Printf("key:%v value:%v", i, getKey(i))
			}(i)
		}
		wg.Wait()
	}
*/
func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			setValue(i, i+100)
			v, _ := s2.Load(i)
			fmt.Printf("key:%v value:%v", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
