package model

type ContainerDeploy struct {
	BaseFields
	Title        string `json:"title"`
	ContainerId       int    `json:"container_id"`
	TeamId       int    `json:"team_id"`
	DeployType       int    `json:"deploy_type"`
	DeployUser       int    `json:"deploy_user"`
	LastAuthorId int    `json:"last_author_id"`
}
