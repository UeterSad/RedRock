package main

//1.该代码运行结果大部分都是9，我认为是循环并发时直接使用i导致的，因为这样写会导致for循环结束后才执行goroutine多线程操作,所以输出的大多都是9
//2.在解决了1后运行程序，果然有了0~9的结果，但是最后发生了死锁，我一看，怎么发送与接收都在main goroutine，于是我把该if语句加入协程来完成它应有的功能
//3.但是出现了i=9那个可能先结束导致main结束而其他还没打出来的情况，我寻思了一下，这time.sleep不也能用吗.jpg（其实最开始想的是wg方法来着，但是我懒XD
//这个程序大概应该没要求按顺序输出吧（
/*func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}*/
import (
	"fmt"
	"time"
)

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n interface{}) {
			fmt.Println(n)
			if n == 9 {
				time.Sleep(100)
				over <- true
			}
		}(i) //把i当作函数参量传进函数中，解决函数不保存i的问题
	}
	<-over
	fmt.Println("over!!!")
}
