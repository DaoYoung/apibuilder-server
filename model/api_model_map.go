package model

type ApiModelMap struct {
	BaseFields
	AuthorId   int    `json:"author_id"`
	ModelId    int    `json:"model_id"`
	TargetId   int    `json:"target_id"`
	TargetType int    `json:"target_type"`
}
