package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ContainerController struct {
	Controller
}
func (this *ContainerController) IsRestRoutePk() bool{
	return true
}
func (action ContainerController) Rester() ControllerInterface {
	actionPtr := &action
	action.Controller.Rester = actionPtr
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Container{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Container{} }
	return  actionPtr
}
func (action *ContainerController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Container).LastAuthorId = user.ID
}