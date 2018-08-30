package helper

import "fmt"

type Speech struct {
	Title string
	Message string
}

func Speak(keyword string, match ...string) Speech {
	var speechChinese = make(map[string]Speech)
	speechChinese["task_appoint"] = Speech{ Title:"任务指定", Message:"%s 指定了任务（%s）给%s"}
	template := speechChinese[keyword]
	if len(match) >0{
		for _, x := range match {
			template.Message = fmt.Sprintf(template.Message, x)
		}
	}
	return template
}
