package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type TaskController struct {
	Controller
}
func (action *TaskController) SetResModel() model.Resource{
	return &(model.Task{})
}
func (action *TaskController) SetResSlice() interface{}{
	return &[]model.Task{}
}
func (action *TaskController) SetSelf() EasyController {
	return action
}

func (action TaskController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource {
		return &(model.Task{})
	}
	actionPtr.GetResSlice = func() interface{} { return &[]model.Task{} }
	actionPtr.Self = actionPtr
	return actionPtr.Controller.DaoService(str)
}
