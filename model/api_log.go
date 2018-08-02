package model

const (
	ApiLogPublish  = 1
	ApiLogCommit   = 2
	ApiLogRender   = 3
	ApiLogTest     = 4
	ApiLogHandOver = 5
	ApiLogRebuild  = 6
	ApiLogNote     = 7
)

type ApiLog struct {
	BaseFields
	UserId     int
	ApiId      int
	FromUserId int
	Type       int
	Status     int
	ForbidUpdateResource
}

func CreateLog(uid int, formUid int, apiId int, logType int, logStatus int) interface{} {
	obj := new(ApiLog)
	obj.UserId = uid
	obj.FromUserId = formUid
	obj.ApiId = apiId
	obj.Type = logType
	obj.Status = logStatus
	return Create(obj)
}
