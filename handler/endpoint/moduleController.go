package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	Controller
}

func (action ModuleController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource { return &(model.Module{}) }
	actionPtr.GetResSlice = func() interface{} { return &[]model.Module{} }
	return actionPtr.Controller.DaoService(str)
}

