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
	api := m.(*model.UserTaskApi).Api()
	tasks := userTask.Depends()
	for _, task := range tasks {
		developer := task.Developer()
		(&model.Notification{}).PoorNew(developer.ID, "task_bind_api", userTask.Title, api.Title)
	}

}