package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
	"log"
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
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := mod.Update(id, &reqInfo)
	c.JSON(http.StatusResetContent, info)
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
