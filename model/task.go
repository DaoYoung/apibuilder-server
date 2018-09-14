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
	AppointUserId int        `json:"appoint_user_id,omitempty"`
	Title         string     `json:"title"`
	Description   string     `json:"description,omitempty"`
	Priority      int        `json:"priority,omitempty"`
	Deadline      *time.Time  `json:"deadline,omitempty"`
	VersionId     string     `json:"version_id,omitempty"`
	HasPrd        int        `json:"has_prd,omitempty"`
	IsCheck       int        `json:"is_check,omitempty"`
	Status        int        `json:"status,omitempty"`
	ExtraTeamTask *[]TeamTask `gorm:"-" json:"team_task,omitempty"`
}

func (mod *Task) ExTeamTasks() {
	teamTasks := &[]TeamTask{}
	FindListWhereKV(teamTasks, "task_id=?", mod.ID, "id", "title", "status")
	if  len(*teamTasks) >0 {
		mod.ExtraTeamTask = teamTasks
	}
}
func (mod Task) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("status")
}