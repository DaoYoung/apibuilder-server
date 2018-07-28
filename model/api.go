package model

import (
	"github.com/jinzhu/gorm"
		)

type Api struct {
	gorm.Model
	TaskId          int    `json:"task_id"`
	ModuleId        string `json:"module_id"`
	Title           string `json:"title"`
	RequestUrl      string `json:"request_url"`
	RequestMethod   string `json:"request_method"`
	RequestMaram    string `json:"request_param"`
	RequestHeader   string `json:"request_header"`
	ResponseContent string `json:"response_content"`
	AuthorId        int    `json:"author_id"`
	Status          int    `json:"status"`
}

func GetApiModel() *BaseFunc {
	bf :=  &BaseFunc{}
	bf.Mod = new(Api)
	bf.ModSlice  = &[]Api{}
	return bf
}
