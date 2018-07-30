package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"reflect"
	"encoding/json"
	"log"
)

//api curd 对象
func ApiAction(str string) func(c *gin.Context) {
	ba := new(BaseAction)
	ba.Mod = model.GetApiModel()
	return CurdAction(ba, str)
}

//Module curd 对象
func ModuleAction(str string) func(c *gin.Context) {
	ba := new(BaseAction)
	ba.Mod = model.GetModuleModel()
	return CurdAction(ba, str)
}

func PublishApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := model.GetApiModel()
	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
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
func CommitApi(c *gin.Context) {
	mod := model.GetApiModel()
	var commitForm model.ApiCommitForm
	err := c.BindJSON(&commitForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT {
		c.JSON(http.StatusForbidden, "Api must published")
	}
	commitAdd(api, &commitForm) //commit LOG AND NOTICE
	info := mod.Update(id, &commitForm)
	c.JSON(http.StatusResetContent, info)
}
func RebuildApi(c *gin.Context) {

}
func NoteApi(c *gin.Context) {

}

func RenderApi(c *gin.Context) {

}
func compareApi(apiOld *model.Api, fieldName string, newVal reflect.Value) (bool, interface{}) {
	v := reflect.ValueOf(*apiOld)
	switch newVal.Kind() {
	case reflect.Int:
		return v.FieldByName(fieldName).Int() == newVal.Int(), v.FieldByName(fieldName).Int()
	case reflect.String:
		return v.FieldByName(fieldName).String() == newVal.String(), v.FieldByName(fieldName).String()
	}
	return true, nil
}
func commitAdd(apiOld *model.Api, comForm *model.ApiCommitForm) (*model.Api, error) {

	v := reflect.ValueOf(*comForm)
	t := reflect.TypeOf(*comForm)
	count := v.NumField()
	chs := make(map[string]interface{})
	for i := 0; i < count; i++ {
		if t.Field(i).Name == "CommitMessage" || t.Field(i).Name == "CommitTaskId" || t.Field(i).Name == "CommitAuthorId" || t.Field(i).Name == "CommitJson" {
			continue
		}
		val := v.Field(i)
		flag, oldval := compareApi(apiOld, t.Field(i).Name, val)
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
		log.Println("no change updated")
		return nil, nil
	}
	commitInfo := new(model.ApiCommit)
	commitInfo.Changes, _ = json.Marshal(chs)
	commitInfo.ApiId = int(apiOld.ID)
	commitInfo.TaskId = comForm.CommitTaskId
	commitInfo.CommitMessage = comForm.CommitMessage
	commitInfo.AuthorId = comForm.CommitAuthorId
	mod := model.GetCommitModel()
	mod.Create(commitInfo)
	model.CreateLog(comForm.CommitAuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_COMMIT, model.API_STATUS_PUBLISH)
	return nil, nil
}
