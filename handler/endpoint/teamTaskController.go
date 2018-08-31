package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
)

type TeamTaskController struct {
	Controller
}
func (action *TeamTaskController) model() model.ResourceInterface {
	return &(model.TeamTask{})
}
func (action *TeamTaskController) modelSlice() interface{} {
	return &[]model.TeamTask{}
}
func (action TeamTaskController) Rester() (actionPtr *TeamTaskController) {
	action.init(&action)
	return  &action
}
func (action *TeamTaskController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.TeamTask).AuthorId = user.ID
}
func (action *TeamTaskController) afterCreate(c *gin.Context, m model.ResourceInterface) {
	task := m.(*model.TeamTask).Task()
	if task.Status == model.StatusInit{
		task = &model.Task{}
		task.Status = model.TaskStatusDispatch
		model.Update(m.(*model.TeamTask).TaskId, task)
	}
	author := model.GetUserFromToken(c)
	team := m.(*model.TeamTask).Team()
	teamLeader := team.Leader()
	(&model.Notification{}).PoorNew( teamLeader.ID, "task_dispatch", author.Username, task.Title, team.TeamName)
}