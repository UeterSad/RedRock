package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("./plan.txt", os.O_CREATE, 0233)
	if err != nil {
		fmt.Println("error:", err)
	}
	wr := bufio.NewWriter(file)                                                            //创建一个写的对象
	wr.WriteString("I’m not afraid of difficulties and insist on learning programming.\n") //存入缓存中
	wr.Flush()                                                                             //将缓存中的内容写到文件
	Rd()
}

func Rd() {
	file, err := os.Open("./plan.txt")
	if err != nil {
		fmt.Println("open file failed:", err)
	}
	defer file.Close()                   //开了别忘记关
	reader := bufio.NewReader(file)      //创建读文件的对象
	line, err := reader.ReadString('\n') //按行读取
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("error:", err)
	}
	fmt.Println(line)
}
