package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	Controller
}
func (action ModuleController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Module{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Module{} }
	return  &action
}
func (action *ModuleController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Module).AuthorId = user.ID
}

