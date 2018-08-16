package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiNoteController struct {
	Controller
}
func (action *ApiNoteController) model() model.ResourceInterface {
	return &(model.ApiNote{})
}
func (action *ApiNoteController) modelSlice() interface{} {
	return &[]model.ApiNote{}
}
func (this *ApiNoteController) parentController() ControllerInterface { return  ApiController{}.Rester()}
func (action ApiNoteController) Rester() *ApiNoteController {
	action.init(&action)
	return  &action
}
func (action *ApiNoteController) RouteName() string {
	return "note"
}
func (action *ApiNoteController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiNote).AuthorId = user.ID
	m.(*model.ApiNote).ApiId,_ = strconv.Atoi(c.Param("api_id"))
}
func (action *ApiNoteController) listCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["api_id"] = c.Param("api_id")
	return condition
}