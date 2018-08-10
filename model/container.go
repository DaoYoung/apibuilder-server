package model

type Container struct {
	BaseFields
	Title        string `json:"title"`
	Status       int    `json:"status"`
	LastAuthorId int    `json:"last_author_id"`
}

