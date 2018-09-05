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
	(&action).InfoFields = []string{"*","Relations()"}
	(&action).ListFields = []string{"id","title","Relations()"}
	return &action
}

func (action *TaskController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Task).AuthorId = user.ID
}
func (this *TaskController) afterUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface) {
	if old.(*model.Task).AppointUserId == 0 && new.(*model.Task).AppointUserId > 0 {
		author := model.GetUserFromToken(c)
		appointUser := &model.User{}
		model.ByID(appointUser, new.(*model.Task).AppointUserId)
		(&model.Notification{}).PoorNew(
			new.(*model.Task).AppointUserId, "task_appoint", author.Username, new.(*model.Task).Title, appointUser.Username)
	}
}
