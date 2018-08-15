package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiNoteController struct {
	Controller
}

func (action ApiNoteController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ApiNote{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ApiNote{} }
	action.Controller.ParentController =  ApiController{}.Rester()
	return  &action
}
func (action *ApiNoteController) RouteName() string {
	return "note"
}
func (action *ApiNoteController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiNote).AuthorId = user.ID
	m.(*model.ApiNote).ApiId,_ = strconv.Atoi(c.Param("api_id"))
}
func (action *ApiNoteController) ListCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["api_id"] = c.Param("api_id")
	return condition
}