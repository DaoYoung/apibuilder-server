package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
		"apibuilder-server/helper"
	"apibuilder-server/app"
	"net/http"
	"strconv"
)

type ContainerParamController struct {
	Controller
}
func (this ContainerParamController) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil{
		panic(err)
	}
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("id")
	obj :=  &[]model.ContainerParam{}
	model.FindListWhereMap(obj, condition, "id asc", page, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}

type ContainerParamList struct {
	Data []model.ContainerParam
}

func (this ContainerParamController) SaveList(c *gin.Context) {
	containerId, _ := strconv.Atoi(c.Param("id"))
	obj :=  &ContainerParamList{}
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	user := model.GetUserFromToken(c)
	for _,param := range (*obj).Data {
		param.LastAuthorId = user.ID
		if param.ID > 0 {
			model.Update(param.ID, &param)
		}else{
			param.ContainerId = containerId
			model.Create(&param)
		}
	}
	helper.ReturnSuccess(c, http.StatusOK, obj)
}