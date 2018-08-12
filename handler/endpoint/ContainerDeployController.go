package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"apibuilder-server/app"
	"apibuilder-server/helper"
	"net/http"
	"strconv"
)

type ContainerDeployController struct {
	Controller
}
func (this ContainerDeployController) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil{
		panic(err)
	}
	condition := make(map[string]interface{})
	condition["container_id"] = c.Param("id")
	obj :=  &[]model.ContainerDeploy{}
	model.FindListWhereMap(obj, condition, "id asc", page, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}
func (this ContainerDeployController) Update(c *gin.Context) {
	obj := &model.ContainerDeploy{}
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("deploy_id"))
	containerId, _ := strconv.Atoi(c.Param("id"))
	condition := make(map[string]interface{})
	condition["container_id"] = containerId
	condition["id"] = id
	user := model.GetUserFromToken(c)
	obj.LastAuthorId = user.ID
	info := model.UpdateWhere(condition, obj)
	helper.ReturnSuccess(c, http.StatusOK, info)
}
func (this ContainerDeployController) Create(c *gin.Context) {
	containerId, _ := strconv.Atoi(c.Param("id"))
	obj := &model.ContainerDeploy{}
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	obj.ContainerId = containerId
	user := model.GetUserFromToken(c)
	obj.LastAuthorId = user.ID
	info := model.Create(obj)
	helper.ReturnSuccess(c, http.StatusCreated, info)
}
