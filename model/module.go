package model

import "apibuilder-server/helper"

type Module struct {
	BaseFields
	Pid      int    `json:"pid"`
	Spid     string `json:"spid"`
	AuthorId int    `json:"author_id"`
	Title    string `json:"title"`
}

func (model Module) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id")
}

func (model Module) ListFields() []string {
	return []string{"id", "title"}
}