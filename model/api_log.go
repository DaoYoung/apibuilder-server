package model

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
	BaseFields
	UserId        int
	ApiId         int
	FromUserId    int
	Type          int
	Status        int
}

func (model *ApiLog) UpdateStruct() interface{} {
	return nil
}

func (model *ApiLog) InitDao() *Dao {
	dao := &Dao{}
	dao.MainResource = model
	dao.SliceResource = &[]ApiLog{}
	return dao
}

func CreateLog(uid int, formUid int , apiId int, logType int, logStatus int) interface{} {
	obj := new(ApiLog)
	obj.UserId = uid
	obj.FromUserId = formUid
	obj.ApiId = apiId
	obj.Type = logType
	obj.Status = logStatus
	return (&(ApiLog{})).InitDao().Create(obj)
}
