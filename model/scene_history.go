package model

type SceneHistory struct {
	BaseFields
	SceneId   int `json:"scene_id"`
	SpendTime int `json:"spend_time"`
	ErrNum    int `json:"err_num"`
	OkNum     int `json:"ok_num"`
	Status    int `json:"status"`
}
