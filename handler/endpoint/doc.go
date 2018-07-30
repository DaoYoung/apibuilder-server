package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"log"
	"reflect"
	"encoding/json"
)

func CreateApi(c *gin.Context) {
	mod := model.GetApiModel()
	var reqInfo model.Api
	err := c.BindJSON(&reqInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	info := mod.Create(&reqInfo)
	c.JSON(http.StatusCreated, info)
}
func UpdateApi(c *gin.Context) {
	mod := model.GetApiModel()
	var reqInfo model.Api
	err := c.BindJSON(&reqInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Println(reqInfo)
	id, _ := strconv.Atoi(c.Param("id"))

	row, err := mod.ByID(id)
	if err != nil {
		c.JSON(http.StatusOK, err)
	}

	api := row.(*model.Api)
	if api.Status == model.API_STATUS_DRAFT && reqInfo.Status == model.API_STATUS_PUBLISH {
		model.CreateLog(api.AuthorId, 0, int(api.ID), model.API_STATUS_PUBLISH, model.API_STATUS_PUBLISH)
		//Publish LOG AND NOTICE
	}

	if api.Status == model.API_STATUS_PUBLISH {
		commitAdd(api, &reqInfo, c) //commit LOG AND NOTICE
	}

	info := mod.Update(id, &reqInfo)

	c.JSON(http.StatusResetContent, info)
}

func commitAdd(apiOld *model.Api, apiNew *model.Api, c *gin.Context) error {
	commitInfo := new(model.ApiCommit)

	vo := reflect.ValueOf(*apiOld)
	v := reflect.ValueOf(*apiNew)
	t := reflect.TypeOf(*apiNew)
	count := v.NumField()
	chs := make(map[string]interface{})

	for i := 0; i < count; i++ {
		val := v.Field(i)
		valo := vo.Field(i)
		if  valo == val || t.Field(i).Name == "CommitMessage" || t.Field(i).Name == "CommitTaskId" || t.Field(i).Name == "CommitAuthorId" || t.Field(i).Name == "CommitJson" {
			continue
		}
		switch val.Kind() {
		case reflect.Int:
			if val.Int() != 0 {
				chint := new(model.CommitChange)
				chint.Before = valo.Int()
				chint.After = val.Int()
				chs[t.Field(i).Name] = *chint
				log.Println(t.Field(i).Name)
			}
		case reflect.String:
			if val.String() != "" {
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
	commitInfo.Changes,_ =  json.Marshal(chs)
	commitInfo.ApiId = int(apiOld.ID)
	commitInfo.TaskId = apiNew.CommitTaskId
	commitInfo.CommitMessage = apiNew.CommitMessage
	commitInfo.AuthorId = apiNew.CommitAuthorId
	mod := model.GetCommitModel()
	mod.Create(commitInfo)
	model.CreateLog(apiNew.CommitAuthorId, 0, int(apiOld.ID), model.APILOG_TYPE_COMMIT, model.API_STATUS_PUBLISH)
	return nil
}

//api curd 对象
func ApiAction(str string) func(c *gin.Context) {
	ba := new(BaseAction)
	ba.Mod = model.GetApiModel()

	return CurdAction(ba, str)
}

//Module curd 对象
func ModuleAction(str string) func(c *gin.Context) {
	log.Print(str)
	ba := new(BaseAction)
	ba.Mod = model.GetModuleModel()

	return CurdAction(ba, str)
}
