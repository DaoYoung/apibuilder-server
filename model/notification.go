package model

import "apibuilder-server/helper"

type Notification struct {
	BaseFields
	UserId     int    `json:"user_id"`
	Type     int    `json:"type"`
	EntityType string `json:"entity_type`
	EntityId   int    `json:"entity_id`
	Title      string `json:"title`
	Message    string `json:"message`
	Status     int    `json:"status`
}

func (mod *Notification) New(userId int, title, message string, entityType string, entityId int)  {
	mod.UserId = userId
	mod.EntityType = entityType
	mod.EntityId = entityId
	mod.Title = title
	mod.Message = message
	Create(mod)
}

func (mod *Notification) PoorNew(userId int, msg helper.Speech)  {
	mod.New(userId,msg.Title,msg.Message,"",0)
}