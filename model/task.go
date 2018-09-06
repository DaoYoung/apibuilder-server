package model

import (
	"time"
	"apibuilder-server/helper"
)

const (
	StatusInit          = 0
	TaskStatusDispatch  = 1
	TaskStatusDevelop   = 2
	TaskStatusTest      = 3
	TaskStatusPublish   = 4
	TaskStatusTerminate = 3
)

type Task struct {
	BaseFields
	AuthorId      int        `json:"author_id"`
	AppointUserId int        `json:"appoint_user_id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Priority      int        `json:"priority"`
	Deadline      time.Time  `json:"deadline"`
	VersionId     string     `json:"version_id"`
	HasPrd        int        `json:"has_prd"`
	IsCheck       int        `json:"is_check"`
	Status        int        `json:"status"`
	ExtraTeamTask *[]TeamTask `gorm:"-" json:"team_task"`
}

func (mod *Task) Relations() {
	teamTasks := &[]TeamTask{}
	FindListWhereKV(teamTasks, "task_id=?", mod.ID, []string{"*"})
	mod.ExtraTeamTask = teamTasks
}
func (mod Task) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("status")
}