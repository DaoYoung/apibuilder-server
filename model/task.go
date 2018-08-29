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
	DependId      int       `json:"depend_id"`
	BindApiId     int       `json:"bind_api_id"`
	Status        int       `json:"status"`
}
