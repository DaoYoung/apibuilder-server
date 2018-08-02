package model

import (
	"apibuilder-server/app"
	"errors"
	"time"
)

type BaseFields struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Resource interface {
	UpdateStruct() interface{}
}

type ForbidUpdateResource struct {

}
func (res *ForbidUpdateResource) UpdateStruct() interface{} {
	return nil
}

func ByID(res Resource, id int) interface{} {

	if err := app.Db.Where("id = ?", id).Last(res).Error; err == nil {
		return res
	} else {
		panic(NotFoundDaoError(errors.New("ByID:" + string(id) + " not found ")))
	}
}

func FindList(res interface{}) interface{} {
	if err := app.Db.Find(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Update(res Resource, id int, data interface{}) interface{} {
	if err := app.Db.Model(res).Where("id = ?", id).Updates(data).Error; err == nil {
		return ByID(res, id)
	} else {
		panic(QueryDaoError(err))
	}
}

func Delete(res Resource, id int) interface{} {
	if err := app.Db.Where("id = ?", id).Delete(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Create(res Resource) interface{} {
	if err := app.Db.Create(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}
