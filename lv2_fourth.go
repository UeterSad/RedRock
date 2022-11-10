package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		t1 := time.Tick(30 * time.Second)
		for range t1 {
			fmt.Println("\n我已经半分钟没听到3G的故事了")
		}
		wg.Done()
	}()
	defer wg.Wait()
	go task1()
	go task2()
	choose()
}
func task1() {
	func() {
		fmt.Println("task1 start...")
		//执行功能
		now := time.Now()
		// 计算下一个时间
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now)) //确定规定时间
		<-t.C
		fmt.Println("\n我还能再战四小时")
		T := time.Tick(24 * time.Hour) //开始每天定时
		for range T {
			fmt.Println("\n我还能再战四小时")
		}
	}()
}
func task2() {
	func() {
		fmt.Println("task2 start...")
		//执行功能
		now := time.Now()
		// 计算下一个时间
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 6, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now)) //确定规定时间
		<-t.C
		fmt.Println("\n我要去图书馆开卷")
		T := time.Tick(24 * time.Hour) //开始每天定时
		for range T {
			fmt.Println("\n我要去图书馆开卷")
		}
	}()
}
func choose() {
	time.Sleep(time.Second * 3)
	var rec string
Loop:
	fmt.Println("是否要新建定时器:Y/N")
	fmt.Scanln(&rec)
	if rec == "Y" {
		fmt.Println("请选择:\n一次性\n可重复")
	Cho:
		fmt.Print("\n输入你的选择:")
		fmt.Scanln(&rec)
		if rec == "一次性" {
			go timer(&rec)
		} else if rec == "可重复" {
			go ticker(&rec)
		} else {
			fmt.Println("\n请输入有效信息")
			goto Cho
		}
	} else if rec == "N" {
		fmt.Println("已退出")
	} else {
		fmt.Println("\n请输入有效信息")
		goto Loop
	}
}
func timer(rec *string) {
	fmt.Println("\n请输入设置时间(示例:2006/1/2-15:04:05)")
	fmt.Scanln(rec)
	t1, err := time.Parse("2006/1/2-15:04:05", *rec)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("\n请输入定时行为")
	fmt.Scanln(rec)
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), now.Location())
	// 检查是否超过当日的时间
	if next.Sub(now) < 0 {
		next = now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), now.Location())
	}
	// 阻塞到执行时间
	t := time.NewTimer(next.Sub(now))
	<-t.C
	fmt.Println(*rec)
}
func ticker(rec *string) {
	fmt.Println("\n请输入设置时间(示例:2006/1/2-15:04:05)")
	fmt.Scanln(rec)
	t1, err := time.Parse("2006/1/2-15:04:05", *rec)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("\n请输入定时行为")
	fmt.Scanln(rec)
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), now.Location())
	// 检查是否超过当日的时间
	if next.Sub(now) < 0 {
		next = now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), now.Location())
	}
	// 阻塞到执行时间
	t := time.NewTimer(next.Sub(now))
	<-t.C
	t2 := time.NewTicker(time.Hour * 24)
	<-t2.C
	fmt.Println(*rec)
}
