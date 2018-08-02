package model

import "apibuilder-server/helper"

type ApiModel struct {
	BaseFields
	AuthorId     int    `json:"author_id"`
	ModelCode       int    `json:"model_code"`
	ModelName   int    `json:"model_name"`
}

func (model *ApiModel) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id")
}

func (model *ApiModel) ListFields() []string {
	return []string{"id", "author_id", "model_code", "model_name"}
}

