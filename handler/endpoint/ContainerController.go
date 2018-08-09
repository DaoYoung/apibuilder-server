package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ContainerController struct {
	Controller
}

func (action ContainerController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action

	actionPtr.Res = &(model.Container{})
	actionPtr.ResSlice = &[]model.Container{}
	actionPtr.TableName = model.Container{}.TableName()
	return actionPtr.Controller.DaoService(str)
}
