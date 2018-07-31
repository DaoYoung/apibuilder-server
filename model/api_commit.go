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

func CreateCommit(chs []byte, msg string, taskId int , apiId int, authorId int) (interface{}, error) {
	commitInfo := new(ApiCommit)
	commitInfo.Changes = chs
	commitInfo.CommitMessage = msg
	commitInfo.ApiId = apiId
	commitInfo.TaskId = taskId
	commitInfo.AuthorId = authorId
	modFunc := GetCommitModel()
	return modFunc.Create(commitInfo)
}

type CommitChange struct {
	Before interface{} `json:"before"`
	After interface{} `json:"after"`
}
type CommitChangeJson struct {
	ChangeJson map[string]CommitChange `json:"change_json"`
}