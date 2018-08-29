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
	Priority       int       `json:"priority"`
	Deadline       time.Time `json:"deadline"`
	Status         int       `json:"status"`
}
