package model

import "apibuilder-server/helper"

type ApiNote struct {
	BaseFields
	ApiId      int            `json:"api_id"`
	ModelId    int            `json:"model_id"`
	AuthorId   int            `json:"author_id"`
	Fkey       string         `json:"fkey"`
	FkeyParent string         `json:"fkey_parent"`
	FkeyToken  string         `json:"fkey_token"`
	Note       string         `json:"note"`
	ModelNotes []ApiModelNote `json:"model_notes" gorm:"-;foreignkey:ModelId"`
}

func (model ApiNote) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id", "api_id")
}
func (model ApiNote) ListFields() []string {
	return []string{"id", "author_id", "fkey", "fkey_parent", "fkey_token", "note", "model_id"}
}
