package model
type Proxy struct {
	BaseFields
	HoldChannelId int    `json:"hold_channel_id"`
	Port int    `json:"port"`
}