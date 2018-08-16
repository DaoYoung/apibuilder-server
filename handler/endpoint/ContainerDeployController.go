package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
				"strconv"
	)

type ContainerDeployController struct {
	Controller
}
func (action *ContainerDeployController) RouteName() string {
	return "deploy"
}
func (action *ContainerDeployController) model() model.ResourceInterface {
	return &(model.ContainerDeploy{})
}
func (action *ContainerDeployController) modelSlice() interface{} {
	return &[]model.ContainerDeploy{}
}
func (action ContainerDeployController) Rester() (actionPtr *ContainerDeployController) {
	action.init(&action)
	return  &action
}
func (this *ContainerDeployController) listCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("container_id")
	return condition
}
func (this *ContainerDeployController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ContainerDeploy).LastAuthorId = user.ID
	m.(*model.ContainerDeploy).ContainerId , _ = strconv.Atoi(c.Param("container_id"))
}

func (this *ContainerDeployController) beforeUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	new.(*model.ContainerDeploy).LastAuthorId = user.ID
	new.(*model.ContainerDeploy).ContainerId , _ = strconv.Atoi(c.Param("container_id"))
}

func (this *ContainerDeployController) updateCondition(c *gin.Context, pk string) map[string]interface{} {
	condition := this.Controller.updateCondition(c, GetRouteID(this))
	condition["container_id"] ,_ = strconv.Atoi(c.Param("container_id"))
	return condition
}