package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ModelNoteController struct {
	Controller
}
func (action *ModelNoteController) model() model.ResourceInterface {
	return &(model.ApiModelNote{})
}
func (action *ModelNoteController) modelSlice() interface{} {
	return &[]model.ApiModelNote{}
}
func (this *ModelNoteController) parentController() ControllerInterface { return  ModelController{}.Rester()}
func (action ModelNoteController) Rester() (*ModelNoteController) {
	action.init(&action)
	return  &action
}
func (action *ModelNoteController) RouteName() string {
	return "note"
}
func (action *ModelNoteController) listCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["model_id"] = c.Param("model_id")
	return condition
}
func (action *ModelNoteController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiModelNote).AuthorId = user.ID
	m.(*model.ApiModelNote).ModelId,_ = strconv.Atoi(c.Param("model_id"))
}