package app

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

var Db *gorm.DB

func InitDb() error {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.Db.User,
		Config.Db.Password,
		Config.Db.Host,
		Config.Db.Port,
		Config.Db.Name,
	)
	if Db, err = gorm.Open("mysql", dsn); err != nil {
		return err
	}
	Db.LogMode(true)
	return nil
}
