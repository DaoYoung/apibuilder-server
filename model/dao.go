package model

import (
	"apibuilder-server/app"
	"errors"
	"time"
	"apibuilder-server/helper"
	"strconv"
	"strings"
	"reflect"
	"fmt"
	"log"
)

//todo view id 客户端注册需要哪些字段，根据场景返回相应字段，避免服务端来关心UI调整
type BaseFields struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (mod BaseFields) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields()
}
func (mod BaseFields) ForbidUpdate() bool {
	return false
}
func ZeroForbidFields(r ResourceInterface) {
	ref := reflect.ValueOf(r).Elem()
	fields := r.ForbidUpdateFields()
	for _, v := range fields {
		if structField := ref.FieldByName(helper.CamelString(v)); structField.IsValid() {
			structField.Set(reflect.Zero(structField.Type()))
		}
	}
}
type ResourceInterface interface {
	ForbidUpdateFields() []string
	ForbidUpdate() bool
}

func filterFuncFields(f *[]string) (r []string) {
	var fSlice  []string
	for _, e := range *f {
		if !strings.Contains(e, "()") {
			fSlice = append(fSlice, e)
			continue
		}
		r = append(r, e)
	}
	log.Print(r,fSlice)
	*f = fSlice
	return r
}
func displayExtraFields(r interface{}, extras []string) {
	v := reflect.ValueOf(r)
	for _, fun := range extras {
		mv := v.MethodByName(strings.Replace(fun, "()", "", -1)) //获取对应的方法
		v.NumMethod()
		if !mv.IsValid() { //判断方法是否存在
			fmt.Println("Func " + fun + " not exist")
			return
		}
		mv.Call(nil)
	}
}
func displayExtraFieldsList(res interface{}, extras []string) {
	var count int
	v := reflect.ValueOf(res)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		count = v.Len()
	}
	for i := 0; i < count; i++ {
		f := v.Index(i).Addr().Interface()
		displayExtraFields(f, extras)
	}
}
func ByID(res ResourceInterface, id int, fields ...string) {
	var funcFields []string
	if len(fields) == 0 {
		fields = append(fields, "*")
	} else {
		funcFields = filterFuncFields(&fields)
	}
	if err := app.Db.Select(fields).Where("id = ?", id).Last(res).Error; err != nil {
		panic(NotFoundDaoError(errors.New("ByID:(" + strconv.Itoa(id) + ") data not found ")))
	} else {
		if funcFields != nil {
			displayExtraFields(res, funcFields)
		}
	}
}

func FindListWhereMap(res interface{}, where map[string]interface{}, order string, page int, limit int, fields ...string) {
	var funcFields []string
	if len(fields) == 0 {
		fields = append(fields, "*")
	} else {
		funcFields = filterFuncFields(&fields)
	}
	offset := limit * (page - 1)
	if err := app.Db.Select(fields).Where(where).Order(order).Offset(offset).Limit(limit).Find(res).Error; err != nil {
		panic(QueryDaoError(err))
	} else {
		if funcFields != nil {
			displayExtraFieldsList(res, funcFields)
		}
	}
}
func FindListWhereKV(res interface{}, whereField string, whereValue interface{}, fields ...string) {
	var funcFields []string
	if len(fields) == 0 {
		fields = append(fields, "*")
	} else {
		funcFields = filterFuncFields(&fields)
	}
	if err := app.Db.Select(fields).Where(whereField, whereValue).Find(res).Error; err != nil {
		panic(QueryDaoError(err))
	} else {
		if funcFields != nil {
			displayExtraFieldsList(res, funcFields)
		}
	}
}

func Find(res ResourceInterface, where map[string]interface{}) ResourceInterface {
	if err := app.Db.Where(where).First(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Update(id int, res ResourceInterface) ResourceInterface {
	if err := app.Db.Model(res).Where("id = ?", id).Updates(res).Error; err == nil {
		ByID(res, id)
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func UpdateWhere(where map[string]interface{}, res ResourceInterface, changes map[string]interface{}) ResourceInterface {
	if err := app.Db.Model(res).Where(where).Updates(changes).Error; err == nil {
		if val, ok := where["id"]; ok {
			ByID(res, val.(int))
			return res
		}
		return Find(res, where)
	} else {
		panic(QueryDaoError(err))
	}
}

func Delete(res ResourceInterface, id int) ResourceInterface {
	if err := app.Db.Where("id = ?", id).Delete(res).Error; err == nil {
		return res
	} else {
		panic(QueryDaoError(err))
	}
}

func Create(res ResourceInterface) ResourceInterface {
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
func ExsitAndFirst(res ResourceInterface) {
	if err := app.Db.Where(res).First(res).Error; err != nil {
		res = nil
	}
}
