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

func (action ModelController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource { return &(model.ApiModel{}) }
	actionPtr.GetResSlice = func() interface{} { return &[]model.ApiModel{} }
	return actionPtr.Controller.DaoService(str)
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

func (action ModelMapController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource { return &(model.ApiModelMap{}) }
	actionPtr.GetResSlice = func() interface{} { return &[]model.ApiModelMap{} }
	return actionPtr.Controller.DaoService(str)
}