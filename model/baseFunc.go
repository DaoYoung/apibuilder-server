package model

import (
	"apibuilder-server/app"
		"errors"
)

type BaseFunc struct {
	Mod interface{}
	ModSlice interface{}
}

func (bf *BaseFunc) ByID(id int) interface{}{
	obj := bf.Mod
	if err := app.Db.Where("id = ?", id).Last(obj).Error; err == nil {
		return obj
	} else {
		panic(NotFoundDaoError(errors.New("ByID:" + string(id) +" not found ")))
	}
}

func (bf *BaseFunc) FindList() interface{} {
	if err := app.Db.Find(bf.ModSlice).Error; err == nil {
		return bf.ModSlice
	} else {
		panic(QueryDaoError(err))
	}
}

func (bf *BaseFunc) Update(id int, data interface{}) interface{} {
	if err := app.Db.Model(bf.Mod).Where("id = ?", id).Updates(data).Error; err == nil {
		return bf.ByID(id)
	} else {
		panic(QueryDaoError(err))
	}
}

func (bf *BaseFunc) Delete(id int) interface{} {
	if err := app.Db.Where("id = ?", id).Delete(bf.Mod).Error; err == nil {
		return bf.Mod
	} else {
		panic(QueryDaoError(err))
	}
}

func (bf *BaseFunc) Create(data interface{}) interface{} {
	if err := app.Db.Create(data).Error; err == nil {
		return data
	} else {
		panic(QueryDaoError(err))
	}
}
