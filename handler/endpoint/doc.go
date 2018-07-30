package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"log"
	"reflect"
	"encoding/json"
	"fmt"
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
	info := mod.Update(id, &reqInfo)
	c.JSON(http.StatusResetContent, info)
}
func RebuildApi(c *gin.Context) {

}
func NoteApi(c *gin.Context) {

}

func RenderApi(c *gin.Context) {

}
func compareApi(apiOld *model.Api, fieldName string, newVal interface{}){
	m := []byte{}
	err := json.Unmarshal(m, apiOld)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
}
func commitAdd(apiOld *model.Api, comForm *model.ApiCommitForm) (*model.Api, error) {
	commitInfo := new(model.ApiCommit)
	v := reflect.ValueOf(*comForm)
	t := reflect.TypeOf(*comForm)
	count := v.NumField()
	chs := make(map[string]interface{})
	for i := 0; i < count; i++ {
		val := v.Field(i)
		if t.Field(i).Name == "CommitMessage" || t.Field(i).Name == "CommitTaskId" || t.Field(i).Name == "CommitAuthorId" || t.Field(i).Name == "CommitJson" {
			continue
		}
		switch val.Kind() {
		case reflect.Int:
			if val.Int() != 0 && valo.Int() != val.Int() {
				chint := new(model.CommitChange)
				chint.Before = valo.Int()
				chint.After = val.Int()
				chs[t.Field(i).Name] = *chint
				log.Println(t.Field(i).Name)
			}
		case reflect.String:
			if val.String() != "" && valo.String() != val.String() {
				chstr := new(model.CommitChange)
				chstr.Before = valo.String()
				chstr.After = val.String()
				chs[t.Field(i).Name] = *chstr
				log.Println(t.Field(i).Name)
			}
		case reflect.Slice: //json被当作slice
			if v.Field(i).Len() != 0 {
				if len(apiNew.CommitJson) > 0 {
					chs[t.Field(i).Name] = apiNew.CommitJson
					log.Println(t.Field(i).Name)
				}

			}
		}
	}
	if len(chs) == 0 {
		return nil
	}
	commitInfo.Changes, _ = json.Marshal(chs)
	commitInfo.ApiId = int(apiOld.ID)
	commitInfo.TaskId = apiNew.CommitTaskId
	commitInfo.CommitMessage = apiNew.CommitMessage
	commitInfo.AuthorId = apiNew.CommitAuthorId
	mod := model.GetCommitModel()
	mod.Create(commitInfo)
	model.CreateLog(apiNew.CommitAuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_COMMIT, model.API_STATUS_PUBLISH)
	return nil
}
