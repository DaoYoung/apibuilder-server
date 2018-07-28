package model

import (
	"apibuilder-server/app"
	"log"
)

type BaseFunc struct {
	Mod interface{}
	ModSlice interface{}
}

func (bf *BaseFunc) ByID(id int) interface{} {
	if bf.Mod == nil {
		return nil
	}


	obj := bf.Mod
	if err := app.Db.Where("id = ?", id).Last(obj).Error; err == nil {
		//log.Print(obj)
		return obj
	} else {
		log.Print(err)
		return nil
	}
}

func (bf *BaseFunc) FindList() interface{} {
	if bf.ModSlice == nil {
		return nil
	}

	if err := app.Db.Find(bf.ModSlice).Error; err == nil {
		//log.Print(obj)
		return bf.ModSlice
	} else {
		log.Print(err)
		return nil
	}
}

func (bf *BaseFunc) update(id int, data map[string]interface{}) interface{} {
	if bf.Mod == nil {
		return nil
	}
	if err := app.Db.Model(bf.Mod).Where("id = ?", id).Updates(data).Error; err == nil {
		return bf.Mod
	} else {
		log.Print(err)
		return nil
	}
}

func (bf *BaseFunc) delete(id int) interface{} {
	if bf.Mod == nil {
		return nil
	}
	if err := app.Db.Where("id = ?", id).Delete(bf.Mod).Error; err == nil {
		return bf.Mod
	} else {
		log.Print(err)
		return nil
	}
}

func (bf *BaseFunc) create(data map[string]interface{}) interface{} {
	if bf.Mod == nil {
		return nil
	}
	if err := app.Db.Create(bf.Mod).Error; err == nil {
		return bf.Mod
	} else {
		log.Print(err)
		return nil
	}
}
