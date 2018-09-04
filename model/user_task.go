package model

import (
	"time"
	)

type UserTask struct {
	BaseFields
	AuthorId     int       `json:"author_id"`
	AppointUserId int       `json:"appoint_user_id"`
	TeamTaskId int       `json:"team_task_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Priority     int       `json:"priority"`
	Deadline     time.Time `json:"deadline"`
	Status       int       `json:"status"`
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
	FindListWhereKV(depends, "task_id=?", mod.ID, []string{"*"})
	tasks := []*UserTask{}
	for _, dep := range *depends {
		tasks = append(tasks, dep.DependTask())
	}
	return tasks
}