package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
	"strconv"
)

func GetApiInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := model.GetApiModel()
	info := mod.ByID(id)
	c.JSON(http.StatusOK, info)
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
	info := mod.Update(reqInfo.ID, &reqInfo)
	c.JSON(http.StatusResetContent, info)
}
func DeleteApi(c *gin.Context) {
	mod := model.GetApiModel()
	id, _ := strconv.Atoi(c.Param("id"))
	mod.Delete(id)
	c.JSON(http.StatusNoContent, gin.H{"id": id})
}
