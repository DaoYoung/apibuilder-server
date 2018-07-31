package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"reflect"
	"errors"
	"encoding/json"
)

func PublishApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := model.GetApiModel()
	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Json content error")
		return
	}
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_PUBLISH {
		c.JSON(http.StatusForbidden, "Api has published")
	} else {
		info,err := mod.Update(id, model.Api{Status: model.API_STATUS_PUBLISH})
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			model.CreateLog(api.AuthorId, 0, int(api.ID), model.APILOG_TYPE_PUBLISH, model.API_STATUS_PUBLISH)
			//todo notice others
			c.JSON(http.StatusOK, info)
		}

	}
}
func CommitApi(c *gin.Context) {
	mod := model.GetApiModel()
	var commitForm model.ApiCommitForm
	err := c.BindJSON(&commitForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT {
		c.JSON(http.StatusForbidden, "Api must published")
	}
	err = commitAdd(api, &commitForm) //commit LOG AND NOTICE
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	info,err := mod.Update(id, &commitForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, info)
	}
}
func RebuildApi(c *gin.Context) {
	mod := model.GetApiModel()
	var apiForm model.Api
	err := c.BindJSON(&apiForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT {
		c.JSON(http.StatusForbidden, "Api must published")
	}
	err = rebuildLog(api, &apiForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	info,err := mod.Update(id, &apiForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, info)
	}
}
func NoteApi(c *gin.Context) {

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
func commitAdd(apiOld *model.Api, comForm *model.ApiCommitForm) error {
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
		return errors.New("no change updated")
	}
	changes, _ := json.Marshal(chs)
	_, err := model.CreateCommit(changes, comForm.CommitMessage, comForm.CommitTaskId , int(apiOld.ID), comForm.CommitAuthorId)
	if err != nil {
		return err
	}
	_, err = model.CreateLog(comForm.CommitAuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_COMMIT, model.API_STATUS_PUBLISH)
	if err != nil {
		return err
	}
	return nil
}
func rebuildLog(apiOld *model.Api, comForm *model.Api)error {
	changes, _ := json.Marshal(apiOld)
	_, err := model.CreateCommit(changes, "rebuild", comForm.TaskId , int(apiOld.ID), comForm.AuthorId)
	if err != nil {
		return err
	}
	_, err = model.CreateLog(comForm.AuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_REBUILD, model.API_STATUS_PUBLISH)
	if err != nil {
		return err
	}
	return nil
}
