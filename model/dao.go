package model

import (
	"apibuilder-server/app"
	"errors"
	"time"
	"apibuilder-server/helper"
	"strconv"
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


func ByID(res Resource, id int) Resource {

	if err := app.Db.Where("id = ?", id).Last(res).Error; err == nil {
		return res
	} else {
		panic(NotFoundDaoError(errors.New("ByID:(" + strconv.Itoa(id) + ") data not found ")))
	}
}

func FindList(res interface{}, where map[string]interface{}) interface{} {
	if err := app.Db.Where(where).Find(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Find(res Resource) Resource {
	if err := app.Db.Where(res).First(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Update(id int, res Resource) Resource {
	if err := app.Db.Model(res).Where("id = ?", id).Updates(res).Error; err == nil {
		return ByID(res, id)
	} else {
		panic(QueryDaoError(err))
	}
}

func Delete(res Resource, id int) Resource {
	if err := app.Db.Where("id = ?", id).Delete(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Create(res Resource) Resource {
	if err := app.Db.Create(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func ExsitAndFirst(res Resource) Resource {
	if err := app.Db.Where(res).First(res).Error; err == nil {
		return res
	} else {
		return nil
	}
}