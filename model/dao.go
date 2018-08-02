package model

import (
	"apibuilder-server/app"
	"errors"
	"time"
	"apibuilder-server/helper"
)

type BaseFields struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (bf BaseFields) ListFields() []string {
	return []string{"*"}
}
func (bf BaseFields) InfoFields() []string {
	return bf.ListFields()
}
func (bf BaseFields) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields()
}

type Resource interface {
	ListFields() []string
	InfoFields() []string
	ForbidUpdateFields() []string
}

type ForbidUpdateResource struct{}

func (bf ForbidUpdateResource) ForbidUpdate() bool {
	return true
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

func Update(id int, res Resource) interface{} {
	if err := app.Db.Where("id = ?", id).Updates(res).Error; err == nil {
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
