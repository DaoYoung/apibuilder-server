package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Controller
}
func (action TaskController) Rester() ControllerInterface {
	actionPtr := &action
	action.Controller.Rester = actionPtr
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Task{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Task{} }
	return  actionPtr
}
func (this *TaskController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Task).AuthorId = user.ID
}
func (this *TaskController) AfterRest(c *gin.Context, m model.ResourceInterface) {
	//todo add task log
}