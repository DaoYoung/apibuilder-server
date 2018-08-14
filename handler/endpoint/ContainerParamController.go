package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
)

type ContainerParamController struct {
	Controller
}
func (action *ContainerParamController) Rester() ControllerInterface {
	action.Controller.Rester = action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ContainerParam{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ContainerParam{} }
	return  action
}
func (this *ContainerParamController) ListCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("id")
	return condition
}
func (this *ContainerParamController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ContainerParam).LastAuthorId = user.ID
	m.(*model.ContainerParam).ContainerId , _ = strconv.Atoi(c.Param("id"))
}