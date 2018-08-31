package model

type UserTaskDepend struct {
	BaseFields
	TaskId   int `json:"task_id"`
	DependId int `json:"depend_id"`
	UserId   int `json:"user_id"`
}
