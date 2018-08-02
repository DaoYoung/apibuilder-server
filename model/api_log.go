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
	UserId     int    `json:"user_id"`
	Type       int    `json:"type"`
	EntityId   int    `json:"entity_id"`
	EntityType string `json:"entity_type"`
	ForbidUpdateResource
}

func CreateLog(uid int, logType int, entityId int, entityType ...string) interface{} {
	obj := new(ApiLog)
	obj.UserId = uid
	obj.Type = logType
	if len(entityType) > 0 {
		obj.EntityType = entityType[0]
	}
	obj.EntityId = entityId
	return Create(obj)
}
