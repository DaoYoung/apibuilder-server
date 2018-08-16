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
func (action *ContainerController) model() model.ResourceInterface {
	return &(model.Container{})
}
func (action *ContainerController) modelSlice() interface{} {
	return &[]model.Container{}
}
func (action ContainerController) Rester() (actionPtr *ContainerController) {
	action.init(&action)
	return  &action
}
func (action *ContainerController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Container).LastAuthorId = user.ID
}