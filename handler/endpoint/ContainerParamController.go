package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
)

type ContainerParamController struct {
	Controller
}
func (action *ContainerParamController) model() model.ResourceInterface {
	return &(model.ContainerParam{})
}
func (action *ContainerParamController) modelSlice() interface{} {
	return &[]model.ContainerParam{}
}
func (action ContainerParamController) Rester() (actionPtr *ContainerParamController) {
	action.init(&action)
	return  &action
}
func (action *ContainerParamController) RouteName() string {
	return "param"
}
func (this *ContainerParamController) listCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("container_id")
	return condition
}
func (this *ContainerParamController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ContainerParam).LastAuthorId = user.ID
	m.(*model.ContainerParam).ContainerId , _ = strconv.Atoi(c.Param("container_id"))
}