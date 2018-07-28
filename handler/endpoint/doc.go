package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apibuilder-server/model"
		"strconv"
)

func GetApiInfo(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))
   	mod := model.GetApiModel()
   	info := mod.ByID(id)
	c.JSON(http.StatusOK, info)
}
func GetApiList(c *gin.Context)  {
	mod := model.GetApiModel()
	list := mod.FindList()
	c.JSON(http.StatusOK, list)
}
func CreateApi(c *gin.Context)  {

	//c.JSON(http.StatusOK, gin.H{"id": id})
}
func UpdateApi(c *gin.Context)  {

	//c.JSON(http.StatusOK, gin.H{"id": id})
}
func DeleteApi(c *gin.Context)  {

	//c.JSON(http.StatusOK, gin.H{"id": id})
}
