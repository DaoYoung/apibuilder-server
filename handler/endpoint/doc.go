package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"log"
			)

func GetApiInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := model.GetApiModel()
	info, err := mod.ByID(id)
	if err == nil {
		//res := info.(*model.Api)
		//var dataContent map[string]interface{}
		//if er := json.Unmarshal([]byte(res.ResponseContent), &dataContent); er == nil {
		//	res.ResponseContentJson = dataContent
		//}
		//var dataParam map[string]interface{}
		//if er := json.Unmarshal([]byte(res.RequestParam), &dataParam); er == nil {
		//	res.RequestParamJson = dataParam
		//}
		//var dataHeader map[string]interface{}
		//if er := json.Unmarshal([]byte(res.RequestHeader), &dataHeader); er == nil {
		//	res.RequestHeaderJson = dataHeader
		//}
		c.JSON(http.StatusOK, info)
	}else {
		c.JSON(http.StatusOK, err)
	}
}
func GetApiList(c *gin.Context) {
	mod := model.GetApiModel()
	list := mod.FindList()
	c.JSON(http.StatusOK, list)
}
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
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := mod.Update(id, &reqInfo)
	c.JSON(http.StatusResetContent, info)
}
func DeleteApi(c *gin.Context) {
	mod := model.GetApiModel()
	id, _ := strconv.Atoi(c.Param("id"))
	mod.Delete(id)
	c.JSON(http.StatusNoContent, gin.H{"id": id})
}

func ApiAction(str string) func(c *gin.Context) {
	ba := new(BaseAction)
	ba.Mod = model.GetApiModel()

	return CurdAction(ba, str)
}

func ModuleAction(str string) func(c *gin.Context) {
	log.Print(str)
	ba := new(BaseAction)
	ba.Mod = model.GetModuleModel()

	return CurdAction(ba, str)
}
