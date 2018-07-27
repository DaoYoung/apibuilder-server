package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var WedDb *gorm.DB

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
	Db.SingularTable(true)

	wedDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.WedDb.User,
		Config.WedDb.Password,
		Config.WedDb.Host,
		Config.WedDb.Port,
		Config.WedDb.Name,
	)
	if WedDb, err = gorm.Open("mysql", wedDsn); err != nil {
		return err
	}
	WedDb.SingularTable(true)

	return nil
}
