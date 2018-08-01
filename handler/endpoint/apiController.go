package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ApiController struct {
	Controller
}

func (action ApiController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.Res = &(model.Api{})
	return actionPtr.Controller.CrudService(str)
}