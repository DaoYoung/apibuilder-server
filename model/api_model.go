package model

import "apibuilder-server/helper"

type ApiModel struct {
	BaseFields
	AuthorId  int    `json:"author_id"`
	ModelCode string `json:"model_code"`
	ModelName string `json:"model_name"`
	ModelNotes []ApiModelNote `gorm:"foreignkey:ModelId;association_foreignkey:ID"`
}

func (model ApiModel) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id")
}

func (model ApiModel) ListFields() []string {
	return []string{"id", "author_id", "model_code", "model_name"}
}
