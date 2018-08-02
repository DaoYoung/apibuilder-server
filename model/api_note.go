package model

import "apibuilder-server/helper"

type ApiNote struct {
	BaseFields
	ApiId      int    `json:"api_id"`
	AuthorId   int    `json:"author_id"`
	Fkey       string `json:"fkey"`
	FkeyParent string `json:"fkey_parent"`
	FkeyToken  string `json:"fkey_token"`
	Note       string `json:"note"`
	ModelId    int    `json:"model_id"`
}

func (model ApiNote) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields("author_id", "api_id")
}

func CreateApiNote(chs []byte, msg string, taskId int, apiId int, authorId int) interface{} {

	return nil
}
