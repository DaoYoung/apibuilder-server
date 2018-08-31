package model

type UserTaskApi struct {
	BaseFields
	TaskId   int `json:"task_id"`
	ApiId int `json:"api_id"`
	UserId   int `json:"user_id"`
}
func (mod *UserTaskApi) UserTask() *UserTask {
	task := &UserTask{}
	ByID(task, mod.TaskId)
	return task
}