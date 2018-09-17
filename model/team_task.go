package model

import "time"

type TeamTask struct {
	BaseFields
	AuthorId       int         `json:"author_id,omitempty"`
	DispatchUserId int         `json:"dispatch_user_id,omitempty"`
	AppointTeamId  int         `json:"appoint_team_id,omitempty"`
	TaskId         int         `json:"task_id,omitempty"`
	Title          string      `json:"title,omitempty"`
	Description    string      `json:"description,omitempty"`
	Deadline       *time.Time  `json:"deadline,omitempty"`
	Status         int         `json:"status,omitempty"`
	ExtraUserTask  *[]UserTask `gorm:"-" json:"user_tasks,omitempty"`
	ExtraTeam      *Team       `gorm:"-" json:"team,omitempty"`
}

func (mod *TeamTask) UserTasks() {
	userTasks := &[]UserTask{}
	FindListWhereKV(userTasks, "team_task_id=?", mod.ID, "id", "title", "appoint_user_id", "status", "Developer()", "Depends()")
	if len(*userTasks) > 0 {
		mod.ExtraUserTask = userTasks
	}
}
func (mod *TeamTask) Task() *Task {
	task := &Task{}
	ByID(task, mod.TaskId)
	return task
}

func (mod *TeamTask) Team() *Team {
	team := &Team{}
	ByID(team, mod.AppointTeamId,"id","team_name")
	mod.ExtraTeam = team
	return team
}
