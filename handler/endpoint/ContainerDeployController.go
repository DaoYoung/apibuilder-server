package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
				"strconv"
	)

type ContainerDeployController struct {
	Controller
}

func (action ContainerDeployController) Rester() ControllerInterface {
	actionPtr := &action
	action.Controller.Rester = actionPtr
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ContainerDeploy{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ContainerDeploy{} }
	return  actionPtr
}
func (this *ContainerDeployController) ListCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("container_id")
	return condition
}
func (this *ContainerDeployController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ContainerDeploy).LastAuthorId = user.ID
	m.(*model.ContainerDeploy).ContainerId , _ = strconv.Atoi(c.Param("container_id"))
}

func (this *ContainerDeployController) BeforeUpdate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ContainerDeploy).LastAuthorId = user.ID
	m.(*model.ContainerDeploy).ContainerId , _ = strconv.Atoi(c.Param("container_id"))
}

func (this *ContainerDeployController) UpdateCondition(c *gin.Context) map[string]interface{} {
	condition := this.Controller.UpdateCondition(c)
	condition["container_id"] ,_ = strconv.Atoi(c.Param("container_id"))
	return condition
}