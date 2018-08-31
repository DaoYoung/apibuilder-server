package model

import "time"

type TeamTask struct {
	BaseFields
	AuthorId       int       `json:"author_id"`
	DispatchUserId int       `json:"dispatch_user_id"`
	AppointTeamId  int       `json:"appoint_team_id"`
	TaskId         int       `json:"task_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Deadline       time.Time `json:"deadline"`
	Status         int       `json:"status"`
}

func (mod *TeamTask) Task() *Task {
	task := &Task{}
	ByID(task, mod.TaskId)
	return task
}

func (mod *TeamTask) Team() *Team {
	team := &Team{}
	ByID(team, mod.AppointTeamId)
	return team
}