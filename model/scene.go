package model
const (
	SceneStatusInit    = 1
	SceneStatusMatch  = 2
	SceneStatusPublish = 3
)
type Scene struct {
	BaseFields
	AuthorId int    `json:"author_id"`
	SerialNo string `json:"serial_no"`
	Title    string `json:"title"`
	RemoteAddr    string `json:"remote_addr"`
	UserAgent    string `json:"user_agent"`
	Status   int    `gorm:"default:1" json:"status"`
}
