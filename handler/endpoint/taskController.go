package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Controller
}
func (action *TaskController) Rester() ControllerInterface {
	action.Controller.Rester = action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Task{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Task{} }
	return  action
}
func (this *TaskController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Task).AuthorId = user.ID
}
func (this *TaskController) AfterRest(c *gin.Context, m model.ResourceInterface) {
	//todo add task log
}