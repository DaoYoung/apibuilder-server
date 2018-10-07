package model

type ProxyReq struct {
	BaseFields
	ProxyChannelId int    `json:"proxy_channel_id"`
	RemoteAddr     string `remote_addr`
	UserAgent      string `user_agent`
	RequestUrl     string `request_url`
	Method         string `method`
	Headers        JSON   `json:"headers"`
	Params         JSON   `json:"params"`
	Response       JSON   `json:"response"`
}
