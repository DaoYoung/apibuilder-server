package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserTaskApiController struct {
	Controller
}
func (action *UserTaskApiController) model() model.ResourceInterface {
	return &(model.UserTaskApi{})
}
func (action *UserTaskApiController) modelSlice() interface{} {
	return &[]model.UserTaskApi{}
}
func (action UserTaskApiController) Rester() (actionPtr *UserTaskApiController) {
	action.init(&action)
	return  &action
}
func (this *UserTaskApiController) parentController() ControllerInterface {
	return  UserTaskController{}.Rester()
}
func (action *UserTaskApiController) RouteName() string {
	return "api"
}

func (action *UserTaskApiController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.UserTaskApi).UserId = user.ID
	m.(*model.UserTaskApi).TaskId,_ = strconv.Atoi(c.Param("usertask_id"))
}
func (action *UserTaskApiController) afterCreate(c *gin.Context, m model.ResourceInterface) {
	userTask := m.(*model.UserTaskApi).UserTask()
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