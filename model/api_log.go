package model

import "github.com/jinzhu/gorm"
const (
	APILOG_TYPE_PUBLISH   int = 1
	APILOG_TYPE_COMMIT int = 2
	APILOG_TYPE_RENDER int = 3
	APILOG_TYPE_TEST int = 4
	APILOG_TYPE_HANDOVER int = 5
	APILOG_TYPE_REBUILD int = 6
	APILOG_TYPE_NOTE int = 7

)
type ApiLog struct {
	gorm.Model
	UserId        int
	ApiId         int
	FromUserId    int
	Type          int
	Status        int
}

func GetLogModel() *BaseFunc {
	bf := &BaseFunc{}
	bf.Mod = new(ApiLog)
	bf.ModSlice = &[]ApiLog{}
	return bf
}

func CreateLog(uid int, formUid int , apiId int, logType int, logStatus int)  {
	obj := new(ApiLog)
	obj.UserId = uid
	obj.FromUserId = formUid
	obj.ApiId = apiId
	obj.Type = logType
	obj.Status = logStatus
	mod := GetLogModel()
	mod.Create(obj)
}
