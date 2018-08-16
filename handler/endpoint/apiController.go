package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
	"reflect"
	"encoding/json"
	"errors"
	"apibuilder-server/app"
	"apibuilder-server/helper"
)

type ApiController struct {
	Controller
}
func (this *ApiController) IsRestRoutePk() bool {
	return true
}
func (action *ApiController) model() model.ResourceInterface {
	return &(model.Api{})
}
func (action *ApiController) modelSlice() interface{} {
	return &[]model.Api{}
}
func (action ApiController) Rester() (actionPtr *ApiController) {
	action.init(&action)
	return  &action
}

func (action *ApiController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Api).AuthorId = user.ID
}

func (action *ApiController) afterUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface) {
	if old.(*model.Api).Status == model.ApiStatusDraft && new.(*model.Api).Status == model.ApiStatusPublish {
		model.CreateLog(old.(*model.Api).AuthorId, model.ApiLogPublish, old.(*model.Api).ID)
	}else{
		commitLog(old.(*model.Api), new.(*model.Api))
	}
}

func commitLog(apiOld *model.Api, comForm *model.Api) {
	v := reflect.ValueOf(*comForm)
	t := reflect.TypeOf(*comForm)
	count := v.NumField()
	chs := make(map[string]interface{})
	for i := 0; i < count; i++ {
		if t.Field(i).Name == "CommitMessage" || t.Field(i).Name == "CommitTaskId" || t.Field(i).Name == "CommitAuthorId" || t.Field(i).Name == "CommitJson" {
			continue
		}
		val := v.Field(i)
		flag, oldval := compareApiData(apiOld, t.Field(i).Name, val)
		switch val.Kind() {
		case reflect.Int:
			if val.Int() != 0 && !flag {
				chint := new(model.CommitChange)
				chint.Before = oldval
				chint.After = val.Int()
				chs[t.Field(i).Name] = *chint
			}
		case reflect.String:
			if val.String() != "" && !flag {
				chstr := new(model.CommitChange)
				chstr.Before = oldval
				chstr.After = val.String()
				chs[t.Field(i).Name] = *chstr
			}
		}
	}
	if len(comForm.CommitHeader) > 0 {
		chs["request_header"] = comForm.CommitHeader
	}
	if len(comForm.CommitParam) > 0 {
		chs["request_param"] = comForm.CommitParam
	}
	if len(comForm.CommitContent) > 0 {
		chs["response_content"] = comForm.CommitContent
	}
	if len(chs) == 0 {
		panic(NOChangeError(errors.New("no change updated")))
	}
	changes, _ := json.Marshal(chs)
	model.CreateCommit(changes, comForm.CommitMessage, comForm.CommitTaskId , apiOld.ID, comForm.CommitAuthorId)
	model.CreateLog(comForm.CommitAuthorId, model.ApiLogCommit, apiOld.ID)

}
func compareApiData(apiOld *model.Api, fieldName string, newVal reflect.Value) (bool, interface{}) {
	v := reflect.ValueOf(*apiOld)
	switch newVal.Kind() {
	case reflect.Int:
		return v.FieldByName(fieldName).Int() == newVal.Int(), v.FieldByName(fieldName).Int()
	case reflect.String:
		return v.FieldByName(fieldName).String() == newVal.String(), v.FieldByName(fieldName).String()
	}
	return true, nil
}



func (this *ApiController)NoteApi(c *gin.Context) {
	var jsonForm model.ApiNote
	var info interface{}
	err := c.BindJSON(&jsonForm)
	if err != nil {
		panic(JsonTypeError(err))
	}
	jsonForm.ApiId, _ = strconv.Atoi(c.Param("id"))
	jsonForm.FkeyToken = jsonForm.FkeyParent + "." +jsonForm.Fkey
	dbNote := model.ApiNote{ApiId: jsonForm.ApiId, FkeyToken: jsonForm.FkeyToken}
	model.ExsitAndFirst(&dbNote)
	if dbNote.ID > 0{
		model.Delete(dbNote, dbNote.ID)
	}
	info = model.Create(&jsonForm)
	helper.ReturnSuccess(c, http.StatusOK, info)
}
func (this *ApiController)NoteApiDetail(c *gin.Context) {
	var resouce model.ApiNote
	apiNotes := &([]model.ApiNote{})
	id, _ := strconv.Atoi(c.Param("id"))
	model.FindListWhereKV(apiNotes, "api_id in (?)", id, resouce.ListFields())
	apiModel := new(model.ApiModel)
	for key,val := range *apiNotes{
		apiModelNotes := &([]model.ApiModelNote{})
		if val.ModelId>0 {
			apiModel.ID = val.ModelId
			app.Db.Model(apiModel).Related(apiModelNotes, "ModelNotes")
			((*apiNotes)[key]).ModelNotes = *apiModelNotes
		}
	}
	helper.ReturnSuccess(c, http.StatusOK, apiNotes)
}
