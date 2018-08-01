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
	AuthorId        int    `json:"author_id"`
	Status          int    `json:"status"`
	ApiCommitForm
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
	CommitParam     JSON   `gorm:"-" json:"commit_param"`
	CommitHeader    JSON   `gorm:"-" json:"commit_header"`
	CommitContent   JSON   `gorm:"-" json:"commit_content"`
	CommitTaskId    int    `gorm:"-" json:"commit_task_id"`
	CommitAuthorId  int    `gorm:"-" json:"commit_author_id"`
}

func (model *Api) UpdateStruct() interface{} {
	return ApiCommitForm{}
}
func (model *Api) InitDao() *Dao {
	dao := &Dao{}
	dao.MainResource = model
	dao.SliceResource = &[]Api{}
	return dao
}


func GetApiModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(Api)
	bf.ModSlice = &[]Api{}
	return bf
}
