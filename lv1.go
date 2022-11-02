package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var ch = make(chan int64)

func add1() {
	for i := 0; i < 50000; i++ {
		x = x + 1
		ch <- 1 //发送1等add2接收，否则阻塞，以此达到交替执行的效果
	} //我的理解是ch接收1后不能再进行add1（ch已满），而add2可以执行，add2执行后同理，因此达到交替效果
	wg.Done()
}
func add2() {
	for i := 0; i < 50000; i++ {
		x = x + 1
		<-ch
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add1() //最开始似乎会竞争导致有时出现99999，知道原理不知道解决方法
	go add2()
	wg.Wait()
	fmt.Println(x)
}
