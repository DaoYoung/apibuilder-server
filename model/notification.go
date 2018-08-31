package model

import "apibuilder-server/helper"

type Notification struct {
	BaseFields
	UserId     int    `json:"user_id"`
	Type       string `json:"type"`
	EntityType string `json:"entity_type`
	EntityId   int    `json:"entity_id`
	Title      string `json:"title`
	Message    string `json:"message`
	Status     int    `json:"status`
}

func (mod *Notification) New(toUserId int, keyword string, speech helper.Speech, entityType string, entityId int) {
	mod.UserId = toUserId
	mod.Type = keyword
	mod.EntityType = entityType
	mod.EntityId = entityId
	mod.Title = speech.Title
	mod.Message = speech.Message
	Create(mod)
}

func (mod *Notification) PoorNew(toUserId int, keyword string, params ...interface{}) {
	mod.New(toUserId, keyword, helper.Speak(keyword, params...), "", 0)
}
