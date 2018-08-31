package helper

import "fmt"

type Speech struct {
	Title string
	Message string
}

func Speak(keyword string, match ...interface{}) Speech {
	var speechChinese = make(map[string]Speech)
	speechChinese["task_appoint"] = Speech{ Title:"任务指定", Message:"%s 指定了任务（%s）给%s"}
	speechChinese["task_dispatch"] = Speech{ Title:"分配团队任务", Message:"%s 分配团队任务（%s）给%s"}
	speechChinese["task_separate"] = Speech{ Title:"分解任务", Message:"%s:%s 分解团队任务（%s） => 开发任务（%s）， 指定给%s"}
	speechChinese["task_develop"] = Speech{ Title:"开发任务", Message:"%s 分配了新开发任务（%s）给你"}
	speechChinese["task_depend"] = Speech{ Title:"设置依赖", Message:"%s设置了 开发任务（%s）依赖 开发任务（%s）"}
	speechChinese["task_bind_api"] = Speech{ Title:"接口文档", Message:"任务（%s）绑定了文档，地址：%s"}
	template := speechChinese[keyword]
	if len(match) >0{
		template.Message = fmt.Sprintf(template.Message, match...)
	}
	return template
}
