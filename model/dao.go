package model

import (
	"apibuilder-server/app"
	"errors"
	"time"
	"apibuilder-server/helper"
	"strconv"
)

//todo view id 客户端注册需要哪些字段，根据场景返回相应字段，避免服务端来关心UI调整
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

func FindListWhereMap(res interface{}, where map[string]interface{}, order string, page int, limit int) {
	offset := limit * (page - 1)
	if err := app.Db.Where(where).Order(order).Offset(offset).Limit(limit).Find(res).Error; err != nil {
		panic(QueryDaoError(err))
	}
}
func FindListWhereKV(res interface{}, whereField string, whereValue interface{}, fields []string) {
	//todo 判断res类型
	if err := app.Db.Select(fields).Where(whereField, whereValue).Find(res).Error; err != nil {
		panic(QueryDaoError(err))
	}
}

func Find(res Resource, where map[string]interface{}) Resource {
	if err := app.Db.Where(where).First(res).Error; err == nil {
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

func CreateNew(tb string, res interface{}) interface{} {
	if err := app.Db.Table(tb).Updates(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}
func ExsitAndFirst(res Resource) {
	if err := app.Db.Where(res).First(res).Error; err != nil {
		res = nil
	}
}
