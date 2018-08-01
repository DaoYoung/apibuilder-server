package model

import "github.com/jinzhu/gorm"

type Module struct {
	gorm.Model
	ModuleUpdate
}

type ModuleUpdate struct {
	Pid      int    `json:"pid"`
	Spid     string `json:"spid"`
	AuthorId int    `json:"author_id"`
	Title    string `json:"title"`
}

func (model *Module) UpdateStruct() interface{} {
	return ModuleUpdate{}
}
func (model *Module) InitDao() *Dao {
	dao := &Dao{}
	dao.MainResource = model
	dao.SliceResource = &[]Module{}
	return dao
}

func GetModuleModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(Module)
	bf.ModSlice = &[]Module{}
	return bf
}
