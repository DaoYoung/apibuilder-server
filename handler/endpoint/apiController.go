package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
	"reflect"
	"encoding/json"
	"errors"
)

type ApiController struct {
	Controller
}

func (action ApiController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.Res = &(model.Api{})
	return actionPtr.Controller.DaoService(str)
}

//todo 提炼valid
func PublishApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := (&(model.Api{})).InitDao()
	row := mod.ByID(id)
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_PUBLISH {
		c.JSON(http.StatusForbidden, "Api has published")
	} else {
		info := mod.Update(id, model.Api{Status: model.API_STATUS_PUBLISH})
		model.CreateLog(api.AuthorId, 0, int(api.ID), model.APILOG_TYPE_PUBLISH, model.API_STATUS_PUBLISH)
		//todo notice others
		c.JSON(http.StatusOK, info)

	}
}

func NoteApi(c *gin.Context) {

}

func CommitApi(c *gin.Context) {
	mod := (&(model.Api{})).InitDao()
	var commitForm model.ApiCommitForm
	err := c.BindJSON(&commitForm)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	row := mod.ByID(id)
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT {
		panic(ForbidError(errors.New("api must published")))
	}
	commitLog(api, &commitForm)
	info := mod.Update(id, &commitForm)
	c.JSON(http.StatusOK, info)
}
func commitLog(apiOld *model.Api, comForm *model.ApiCommitForm) {
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
	model.CreateCommit(changes, comForm.CommitMessage, comForm.CommitTaskId , int(apiOld.ID), comForm.CommitAuthorId)
	model.CreateLog(comForm.CommitAuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_COMMIT, model.API_STATUS_PUBLISH)

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
func RebuildApi(c *gin.Context) {
	mod := (&(model.Api{})).InitDao()
	var apiForm model.Api
	err := c.BindJSON(&apiForm)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	row := mod.ByID(id)
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT {
		panic(ForbidError(errors.New("api must published")))
	}
	rebuildLog(api, &apiForm)
	info := mod.Update(id, &apiForm)
	c.JSON(http.StatusOK, info)
}
func rebuildLog(apiOld *model.Api, comForm *model.Api) {
	changes, _ := json.Marshal(apiOld)
	model.CreateCommit(changes, "rebuild", comForm.TaskId , int(apiOld.ID), comForm.AuthorId)
	model.CreateLog(comForm.AuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_REBUILD, model.API_STATUS_PUBLISH)
}