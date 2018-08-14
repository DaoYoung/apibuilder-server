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
func (action *ModelController) GetRestModel() model.ResourceInterface{
	return &(model.ApiModel{})
}
func (action *ModelController) GetRestModelSlice() interface{}{
	return &[]model.ApiModel{}
}
func (action *ModelController) GetRester() *ModelController {
	return action
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
func (action *ModelMapController) GetRestModel() model.ResourceInterface{
	return &(model.ApiModelMap{})
}
func (action *ModelMapController) GetRestModelSlice() interface{}{
	return &[]model.ApiModelMap{}
}
func (action *ModelMapController) GetRester() *ModelMapController {
	return action
}