package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Controller
}
func (action *TaskController) model() model.ResourceInterface {
	return &(model.Task{})
}
func (action *TaskController) modelSlice() interface{} {
	return &[]model.Task{}
}
func (action TaskController) Rester() (actionPtr *TaskController) {
	action.init(&action)
	return  &action
}

func (action *TaskController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Task).AuthorId = user.ID
}
func (action *TaskController) assign(c *gin.Context) {
	//todo add task log
}
