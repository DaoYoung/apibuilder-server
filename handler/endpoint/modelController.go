package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
)

type ModelController struct {
	Controller
}

func (action ModelController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.Res = &(model.ApiModel{})
	actionPtr.ResSlice = &[]model.ApiModel{}
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
	cloneNote := model.ApiModelNote{ModelId: jsonForm.ModelId, ParentId: jsonForm.ParentId, ModelKey: jsonForm.ModelKey}
	dbData := model.ExsitAndFirst(&cloneNote)
	if dbData != nil {
		dbNote := dbData.(*model.ApiModelNote)
		model.Delete(dbNote, dbNote.ID)
	}
	info = model.Create(&jsonForm)
	ReturnSuccess(c, http.StatusOK, info)
}

func NoteModelDetail(c *gin.Context) {
	condition := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Param("id"))
	condition["model_id"] = id
	modelNotes := &([]model.ApiModelNote{})
	model.FindListWhereMap(modelNotes, condition)
	ReturnSuccess(c, http.StatusOK, modelNotes)
}

type ModelMapController struct {
	Controller
}

func (action ModelMapController) CrudService(str string) func(c *gin.Context) {
	actionPtr := &action
	actionPtr.Res = &(model.ApiModelMap{})
	actionPtr.ResSlice = &[]model.ApiModelMap{}
	return actionPtr.Controller.DaoService(str)
}