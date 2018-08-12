package model

type ContainerParam struct {
	BaseFields
	ContainerId  int    `json:"container_id"`
	KeyString    string `json:"key_string"`
	ValueString  string `json:"value_string"`
	LastAuthorId int    `json:"last_author_id"`
}
