package model

import (
	"github.com/jinzhu/gorm"
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

func GetApiModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(Api)
	bf.ModSlice = &[]Api{}
	return bf
}
