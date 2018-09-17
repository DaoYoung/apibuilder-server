package model

type UserTaskDepend struct {
	BaseFields
	TaskId   int `json:"task_id,omitempty"`
	DependId int `json:"depend_id,omitempty"`
	UserId   int `json:"user_id,omitempty"`
}
func (mod *UserTaskDepend) UserTask() *UserTask {
	task := &UserTask{}
	ByID(task, mod.TaskId)
	return task
}
func (mod *UserTaskDepend) DependTask() *UserTask {
	task := &UserTask{}
	ByID(task, mod.DependId)
	return task
}