package model

import (
	"github.com/jinzhu/gorm"
)

type ApiCommit struct {
	gorm.Model
	TaskId        int
	ApiId         int
	AuthorId      int
	CommitMessage string
	Changes       JSON
}

func GetCommitModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(ApiCommit)
	bf.ModSlice = &[]ApiCommit{}
	return bf
}



type CommitChange struct {
	Before interface{} `json:"before"`
	After interface{} `json:"after"`
}
type CommitChangeJson struct {
	ChangeJson map[string]CommitChange `json:"change_json"`
}