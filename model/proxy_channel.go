package model

type ProxyChannel struct {
	BaseFields
	AuthorId int    `json:"author_id"`
	ProxyId     int `json:"proxy_id"`
	Status   int    `gorm:"default:1" json:"status"`
}
