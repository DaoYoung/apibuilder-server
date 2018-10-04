package model

type SceneReq struct {
	BaseFields
	Url      string `json:"url"`
	Method   string `json:"method"`
	Params   JSON   `json:"params,omitempty"`
	Response JSON   `json:"response,omitempty"`
}
