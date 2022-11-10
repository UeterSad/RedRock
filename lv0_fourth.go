package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var chatline string
	fmt.Println("请输入你的对话内容:") //我这里用fuckyou和hello演示（毕竟我没有那么多脏话数据XD
	fmt.Scanln(&chatline)
	fmt.Println() //换个行好看一点
	fmt.Println("拒绝输出型:")
	if strings.Contains(chatline, "fuck") {
		fmt.Println("根据相关法律法规，不得发布反动、暴力、色情或辱骂他人的内容！") /*
			因为strings包有Contains函数且返回bool值，可以配合if来实现检测字符串内容并做出相应措施
			可以做出拒绝输出效果
		*/
	} else {
		fmt.Println(chatline, time.Now()) //反之正常输出，并且带上发送时间
	}
	fmt.Println()
	fmt.Println("屏蔽词型:")
	if strings.Contains(chatline, "fuck") {
		word := strings.Replace(chatline, "fuck", "**", 1)
		fmt.Println(word, time.Now()) //广大网站的屏蔽词效果
	} else {
		fmt.Println(chatline, time.Now())
	}
}
