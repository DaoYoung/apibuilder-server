package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ContainerController struct {
	Controller
}
func (action *ContainerController) Rester() ControllerInterface {
	action.Controller.Rester = action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Container{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Container{} }
	return  action
}
func (action *ContainerController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Container).LastAuthorId = user.ID
}