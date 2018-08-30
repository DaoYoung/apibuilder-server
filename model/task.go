package model

import "time"

type Task struct {
	BaseFields
	AuthorId      int       `json:"author_id"`
	AppointUserId int       `json:"appoint_user_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Priority      int       `json:"priority"`
	Deadline      time.Time `json:"deadline"`
	VersionTag   string    `json:"version_tag"`
	HasPrd       int       `json:"has_prd"`
	IsCheck      int       `json:"is_check"`
	Status        int       `json:"status"`
}
