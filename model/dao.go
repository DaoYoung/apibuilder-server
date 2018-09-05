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
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (mod BaseFields) ListFields() []string {
	return []string{"*"}
}
func (mod BaseFields) InfoFields() []string {
	return mod.ListFields()
}
func (mod BaseFields) ForbidUpdateFields() []string {
	return helper.SetForbidUpdateFields()
}


type ResourceInterface interface {
	ListFields() []string
	InfoFields() []string
	ForbidUpdateFields() []string

}

type ForbidUpdateResource struct{}

func (bf ForbidUpdateResource) ForbidUpdate() bool {
	return true
}
func filterFuncFields(f *[]string) (r []string){
	fSlice := *f
	fLen := len(fSlice)-1
	for i,e := range fSlice{
		if !strings.Contains(e, "()"){
			continue
		}
		r = append(r, e)

		if fLen == i{
			fSlice = fSlice[:i]
			log.Print(i, e, fSlice)
		}else {
			log.Print(i, e)
			fSlice = append(fSlice[:i], fSlice[i+1:]...)
		}
	}
	*f = fSlice
	return  r
}
func DisplayExtraFields(r ResourceInterface, f []string){
	v := reflect.ValueOf(r)
	//if v.Kind() == reflect.Ptr {
	//	v = v.Elem()
	//}
	//log.Print(f)
	for _,fun := range f{
		mv := v.MethodByName(strings.Replace(fun, "()", "", -1)) //获取对应的方法
		v.NumMethod()
		if !mv.IsValid() {            //判断方法是否存在
			fmt.Println("Func " + fun +" not exist")
			return
		}
		mv.Call(nil)
	}
}
func ByID(res ResourceInterface, id int, field ...string) {
	var funcFields []string
	if len(field) == 0{
		field = append(field, "*")
	}else{
		funcFields = filterFuncFields(&field)
		log.Print(funcFields, &field)
	}
	if err := app.Db.Select(field).Where("id = ?", id).Last(res).Error; err != nil {
		panic(NotFoundDaoError(errors.New("ByID:(" + strconv.Itoa(id) + ") data not found ")))
	}else{
		if funcFields != nil {
			DisplayExtraFields(res, funcFields)
		}
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

func UpdateWhere(where map[string]interface{}, res ResourceInterface) ResourceInterface {
	if err := app.Db.Model(res).Where(where).Updates(res).Error; err == nil {
		if val, ok := where["id"]; ok {
			ByID(res, val.(int))
			return res
		}
		return res
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
