package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
	"apibuilder-server/helper"
	"apibuilder-server/app"
)

type ModelController struct {
	Controller
}
func (action ModelController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ApiModel{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ApiModel{} }
	return  &action
}
func (action *ModelController) BeforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiModel).AuthorId = user.ID
}
func (action *ModelController) IsRestRoutePk() bool{
	return true
}
func (action *ModelController) RouteName() string {
	return "model"
}
func NoteModel(c *gin.Context) {
	var jsonForm model.ApiModelNote
	var info interface{}
	err := c.BindJSON(&jsonForm)
	if err != nil {
		panic(JsonTypeError(err))
	}
	jsonForm.ModelId, _ = strconv.Atoi(c.Param("id"))
	dbNote := model.ApiModelNote{ModelId: jsonForm.ModelId, ParentId: jsonForm.ParentId, ModelKey: jsonForm.ModelKey}
	model.ExsitAndFirst(&dbNote)
	if dbNote.ID >0 {
		model.Delete(dbNote, dbNote.ID)
	}
	info = model.Create(&jsonForm)
	helper.ReturnSuccess(c, http.StatusOK, info)
}

func NoteModelDetail(c *gin.Context) {
	condition := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Param("id"))
	condition["model_id"] = id
	modelNotes := &([]model.ApiModelNote{})
	model.FindListWhereMap(modelNotes, condition, "", 1, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, modelNotes)
}

type ModelMapController struct {
	Controller
}
func (action ModelMapController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.ApiModelMap{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.ApiModelMap{} }
	action.Controller.ParentController = ModelController{}.Rester()
	return  &action
}
func (action *ModelMapController) RouteName() string {
	return "map"
}
func (action *ModelMapController) BeforeRest(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ApiModelMap).AuthorId = user.ID
}
