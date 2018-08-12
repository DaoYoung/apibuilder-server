package model

type ContainerDeploy struct {
	BaseFields
	ContainerId       int    `json:"container_id"`
	TeamId       int    `json:"team_id"`
	DeployType       int    `json:"deploy_type"`
	DeployUser       int    `json:"deploy_user"`
	LastAuthorId int    `json:"last_author_id"`
}
