package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
	"errors"
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
		if dbNote.AuthorId != jsonForm.AuthorId {
			panic(ForbidError(errors.New("you can't post note")))
		}
		info = model.Update(dbNote.ID, &jsonForm)
	} else {
		info = model.Create(&jsonForm)
	}
	ReturnSuccess(c, http.StatusOK, info)
}

func NoteModelDetail(c *gin.Context) {
	condition := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Param("id"))
	condition["model_id"] = id
	ReturnSuccess(c, http.StatusOK, model.FindList(&([]model.ApiModelNote{}), condition))
}
