package model

import "github.com/jinzhu/gorm"

type Module struct {
	gorm.Model
	Pid      int    `json:"pid"`
	AuthorId int    `json:"author_id"`
	Title    string `json:"title"`
}

func GetModuleModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(Module)
	bf.ModSlice = &[]Module{}
	return bf
}
