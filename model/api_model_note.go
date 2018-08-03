package model

import (
	"apibuilder-server/helper"
)

type ApiModelNote struct {
	BaseFields
	AuthorId int    `json:"author_id"`
	ModelId  int    `json:"model_id"`
	ModelKey string `json:"model_key"`
	ParentId int    `json:"parent_id"`
	Note     string `json:"note"`
}

func (model ApiModelNote) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id")
}