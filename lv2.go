package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int) //用无缓存的channel实现交替打印
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- 1
			//奇数
			if i%2 == 1 {
				fmt.Println("goroutine1: ", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			//偶数
			if i%2 == 0 {
				fmt.Println("goroutine2: ", i)
			}
		}
	}()
	wg.Wait()
}
