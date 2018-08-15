package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ModelNoteController struct {
	Controller
}

func (action ModelNoteController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ApiModelNote{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ApiModelNote{} }
	action.Controller.ParentController =  ModelController{}.Rester()
	return  &action
}
func (action *ModelNoteController) RouteName() string {
	return "note"
}
func (action *ModelNoteController) ListCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["model_id"] = c.Param("model_id")
	return condition
}
func (action *ModelNoteController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiModelNote).AuthorId = user.ID
	m.(*model.ApiModelNote).ModelId,_ = strconv.Atoi(c.Param("model_id"))
}