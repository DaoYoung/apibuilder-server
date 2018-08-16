package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	Controller
}
func (action *ModuleController) model() model.ResourceInterface {
	return &(model.Module{})
}
func (action *ModuleController) modelSlice() interface{} {
	return &[]model.Module{}
}
func (action ModuleController) Rester() (actionPtr *ModuleController) {
	action.init(&action)
	return  &action
}
func (action *ModuleController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Module).AuthorId = user.ID
}

