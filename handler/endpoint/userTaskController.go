package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type UserTaskController struct {
	Controller
}
func (action *UserTaskController) IsRestRoutePk() bool {
	return true
}
func (action *UserTaskController) model() model.ResourceInterface {
	return &(model.UserTask{})
}
func (action *UserTaskController) modelSlice() interface{} {
	return &[]model.UserTask{}
}
func (action UserTaskController) Rester() (actionPtr *UserTaskController) {
	action.init(&action)
	return  &action
}
func (action *UserTaskController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.UserTask).AuthorId = user.ID
}
func (action *UserTaskController) afterCreate(c *gin.Context, m model.ResourceInterface) {
	teamTask := m.(*model.UserTask).TeamTask()
	task := teamTask.Task()
	if task.Status == model.TaskStatusDispatch{
		task = &model.Task{}
		task.Status = model.TaskStatusDevelop
		model.Update(teamTask.TaskId, task)
	}
	author := model.GetUserFromToken(c)
	developer := m.(*model.UserTask).Developer()
	(&model.Notification{}).PoorNew( task.AppointUserId, "task_separate", author.Team().TeamName,author.Username,teamTask.Title,m.(*model.UserTask).Title, developer.Username)

	(&model.Notification{}).PoorNew( m.(*model.UserTask).AppointUserId, "task_develop",author.Username,m.(*model.UserTask).Title)
}
