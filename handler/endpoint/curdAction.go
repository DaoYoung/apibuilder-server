package endpoint
import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type ApiAction struct {
	BaseAction
}

func (action ApiAction) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.ModFunc = model.GetApiModel()
	return actionPtr.BaseAction.CrudService(str)
}

type ModuleAction struct {
	BaseAction
}

func (action ModuleAction) CrudService(str string) func(c *gin.Context)  {
	//action.CrudService()
	actionPtr := &action
	actionPtr.ModFunc = model.GetModuleModel()
	return actionPtr.BaseAction.CrudService(str)
}
