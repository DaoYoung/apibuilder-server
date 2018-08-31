package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)


type UserTaskDependController struct {
	Controller
}
func (action *UserTaskDependController) model() model.ResourceInterface {
	return &(model.UserTaskDepend{})
}
func (action *UserTaskDependController) modelSlice() interface{} {
	return &[]model.UserTaskDepend{}
}
func (action UserTaskDependController) Rester() (actionPtr *UserTaskDependController) {
	action.init(&action)
	return  &action
}
func (this *UserTaskDependController) parentController() ControllerInterface {
	return  UserTaskController{}.Rester()
}
func (action *UserTaskDependController) RouteName() string {
	return "depend"
}

func (action *UserTaskDependController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.UserTaskDepend).UserId = user.ID
	m.(*model.UserTaskDepend).TaskId,_ = strconv.Atoi(c.Param("usertask_id"))
}