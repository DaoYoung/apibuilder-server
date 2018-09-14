package model

import (
	"time"
	)

type UserTask struct {
	BaseFields
	AuthorId     int       `json:"author_id,omitempty"`
	AppointUserId int       `json:"appoint_user_id,omitempty"`
	TeamTaskId int       `json:"team_task_id,omitempty"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	Priority     int       `json:"priority,omitempty"`
	Deadline     *time.Time `json:"deadline,omitempty"`
	Status       int       `json:"status,omitempty"`
}
func (mod *UserTask) TeamTask() *TeamTask {
	task := &TeamTask{}
	ByID(task, mod.TeamTaskId)
	return task
}
func (mod *UserTask) Developer() *User {
	user := &User{}
	ByID(user, mod.AppointUserId)
	return user
}
func (mod *UserTask) Depends() []*UserTask {
	depends := &[]UserTaskDepend{}
	FindListWhereKV(depends, "task_id=?", mod.ID)
	tasks := []*UserTask{}
	for _, dep := range *depends {
		tasks = append(tasks, dep.DependTask())
	}
	return tasks
}