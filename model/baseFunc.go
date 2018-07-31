package model

import (
	"apibuilder-server/app"
		"errors"
)

type BaseFunc struct {
	Mod interface{}
	ModSlice interface{}
}

func (bf *BaseFunc) ByID(id int) (interface{}, error){
	if bf.Mod == nil {
		return nil, errors.New("model not defined")
	}
	obj := bf.Mod
	if err := app.Db.Where("id = ?", id).Last(obj).Error; err == nil {
		return obj, nil
	} else {
		return nil, err
	}
}

func (bf *BaseFunc) FindList() (interface{}, error) {
	if bf.ModSlice == nil {
		return nil, errors.New("model not defined")
	}

	if err := app.Db.Find(bf.ModSlice).Error; err == nil {
		//log.Print(obj)
		return bf.ModSlice, nil
	} else {
		return nil, err
	}
}

func (bf *BaseFunc) Update(id int, data interface{}) (interface{}, error) {
	if bf.Mod == nil {
		return nil, errors.New("model not defined")
	}
	if err := app.Db.Model(bf.Mod).Where("id = ?", id).Updates(data).Error; err == nil {
		return bf.Mod, nil
	} else {
		return nil, err
	}
}

func (bf *BaseFunc) Delete(id int) (interface{}, error) {
	if bf.Mod == nil {
		return nil, errors.New("model not defined")
	}
	if err := app.Db.Where("id = ?", id).Delete(bf.Mod).Error; err == nil {
		return bf.Mod, nil
	} else {
		return nil, err
	}
}

func (bf *BaseFunc) Create(data interface{}) (interface{}, error) {
	if bf.Mod == nil {
		return nil, errors.New("model not defined")
	}
	if err := app.Db.Create(data).Error; err == nil {
		return data, nil
	} else {
		return nil, err
	}
}
