package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("开始演示源氏技能...")
	skill1 := func(s string) {
		fmt.Println("竜神の剣を喰らえ!", s)
	}
	skill2 := func(s string) {
		fmt.Println("フッ、来るがいい", s)
	}
	skill3 := func(s string) {
		fmt.Println("sei", s)
	}
	var skillname string
	fmt.Println("可使用技能:\n斩\n影\n闪")
	fmt.Println()
	fmt.Println("请输入你想使用的技能名字:")
	fmt.Scanln(&skillname)
	fmt.Println()
	switch skillname {
	case "斩":
		ReleaseSkill("斩", skill1)
	case "影":
		ReleaseSkill("影", skill2)
	case "闪":
		ReleaseSkill("闪", skill3)
	}
	//我是分界线
	fmt.Println()
	var skill map[string]string
	skill = make(map[string]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Println("是否要自定义技能？Y/N")
		var rec string
		fmt.Scanln(&rec)
		if rec == "Y" {
			fmt.Println("请输入技能名:")
			fmt.Scanln(&skillname)
			if Banword(skillname) {
				continue
			}
			fmt.Println("请输入技能描述:")
			fmt.Scanln(&rec)
			if Banword(rec) {
				continue
			}
			skill[skillname] = rec
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Println("释放技能...")
	fmt.Println("请输入要释放的技能名:")
	fmt.Scanln(&skillname)
	fmt.Println("释放中...")
	ReleaseSkill(skill[skillname], Releasefunc())
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func Banword(s string) bool {
	if strings.Contains(s, "fuck") {
		fmt.Println("根据相关法律法规，不得发布反动、暴力、色情或辱骂他人的内容！")
		return true
	} else if strings.Contains(s, "傻逼") {
		fmt.Println("根据相关法律法规，不得发布反动、暴力、色情或辱骂他人的内容！")
		return true
	} else if strings.Contains(s, "他妈的") {
		fmt.Println("根据相关法律法规，不得发布反动、暴力、色情或辱骂他人的内容！")
		return true
	}
	return false
}
func Releasefunc() func(s string) {
	return func(s string) {
		fmt.Println(s)
	}
}
