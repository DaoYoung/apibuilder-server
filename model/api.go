package model

import (
	"github.com/jinzhu/gorm"
)

const (
	API_STATUS_DRAFT   int = 0
	API_STATUS_PUBLISH int = 1
)

type Api struct {
	gorm.Model
	TaskId          int    `json:"task_id"`
	ModuleId        int    `json:"module_id"`
	AuthorId        int    `json:"author_id"`
	Title           string `json:"title"`
	RequestUrl      string `json:"request_url"`
	RequestMethod   string `json:"request_method"`
	RequestParam    JSON   `json:"request_param"`
	RequestHeader   JSON   `json:"request_header"`
	ResponseContent JSON   `json:"response_content"`
	Status          int    `json:"status"`
}

type ApiCommitForm struct {
	TaskId          int    `json:"task_id"`
	ModuleId        int    `json:"module_id"`
	Title           string `json:"title"`
	RequestUrl      string `json:"request_url"`
	RequestMethod   string `json:"request_method"`
	RequestParam    JSON   `json:"request_param"`
	RequestHeader   JSON   `json:"request_header"`
	ResponseContent JSON   `json:"response_content"`
	CommitMessage   string `gorm:"-" json:"commit_message"`
	CommitJson   JSON `gorm:"-" json:"commit_json"`
	CommitTaskId    int    `gorm:"-" json:"commit_task_id"`
	CommitAuthorId        int    `gorm:"-" json:"commit_author_id"`
}

func GetApiModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(Api)
	bf.ModSlice = &[]Api{}
	return bf
}
