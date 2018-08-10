package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ContainerController struct {
	Controller
}

func (action ContainerController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource { return &(model.Container{}) }
	actionPtr.GetResSlice = func() interface{} { return &[]model.Container{} }
	return actionPtr.Controller.DaoService(str)
}
