package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ModelController struct {
	Controller
}

func (action ModelController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.Res = &(model.ApiModel{})
	actionPtr.ResSlice = &[]model.ApiModel{}
	return actionPtr.Controller.DaoService(str)
}
