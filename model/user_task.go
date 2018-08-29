package model

import "time"

type UserTask struct {
	BaseFields
	AuthorId     int       `json:"author_id"`
	AssignUserId int       `json:"assign_user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Priority     int       `json:"priority"`
	Deadline     time.Time `json:"deadline"`
	VersionTag   string    `json:"version_tag"`
	HasPrd       int       `json:"has_prd"`
	IsCheck      int       `json:"is_check"`
	Status       int       `json:"status"`
}
