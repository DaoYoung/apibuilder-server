package model

import (
	"apibuilder-server/app"
	"errors"
	"log"
)

type Resource interface {
	InitDao() *Dao
	UpdateStruct() interface{}
}

type Dao struct {
	MainResource Resource
	SliceResource interface{}
}

func (res *Dao) UpdateStruct() interface{} {
	return nil
}

func (res *Dao) ByID(id int) interface{} {

	if err := app.Db.Where("id = ?", id).Last(res.MainResource).Error; err == nil {
		log.Println(res.MainResource)
		return res.MainResource
	} else {
		panic(NotFoundDaoError(errors.New("ByID:" + string(id) + " not found ")))
	}
}

func (res *Dao) FindList() interface{} {
	if err := app.Db.Find(res.SliceResource).Error; err == nil {
		return res.SliceResource
	} else {
		panic(QueryDaoError(err))
	}
}

func (res *Dao) Update(id int, data interface{}) interface{} {
	if err := app.Db.Model(res.MainResource).Where("id = ?", id).Updates(data).Error; err == nil {
		return res.ByID(id)
	} else {
		panic(QueryDaoError(err))
	}
}

func (res *Dao) Delete(id int) interface{} {
	if err := app.Db.Where("id = ?", id).Delete(res.MainResource).Error; err == nil {
		return res.MainResource
	} else {
		panic(QueryDaoError(err))
	}
}

func (res *Dao) Create(data interface{}) interface{} {
	if err := app.Db.Create(data).Error; err == nil {
		return data
	} else {
		panic(QueryDaoError(err))
	}
}
