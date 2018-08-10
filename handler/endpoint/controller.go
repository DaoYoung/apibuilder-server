package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"errors"
	"apibuilder-server/helper"
	"apibuilder-server/app"
)

type ControllerInterface interface {
	CrudService(str string) func(c *gin.Context)
}

type Controller struct {
	GetResModel func()  model.Resource
	GetResSlice func()  interface{}//https://golang.org/doc/faq#convert_slice_of_interface
}

func (this *Controller) CrudService(str string) func(c *gin.Context) {
	panic(ForbidError(errors.New("not support model curd")))
}

func (this *Controller) Info(c *gin.Context) {
	obj := this.GetResModel()
	id, _ := strconv.Atoi(c.Param("id"))
	info := model.ByID(obj, id)
	helper.ReturnSuccess(c, http.StatusOK, info)
}
func (this *Controller) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil{
		panic(err)
	}
	obj := this.GetResSlice()
	condition := helper.MapUrlQuery(c.Request.URL.Query(), this.GetResModel())
	model.FindListWhereMap(obj, condition, "id desc", page, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}
func (this *Controller) BeforeCreate(c *gin.Context, m model.Resource) {

}
func (this *Controller) Create(c *gin.Context) {
	obj := this.GetResModel()
	this.be
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	info := model.Create(obj)
	helper.ReturnSuccess(c, http.StatusCreated, info)
}

func (this *Controller) Update(c *gin.Context) {
	obj := this.GetResModel()
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := model.Update(id, obj)
	helper.ReturnSuccess(c, http.StatusOK, info)
}

func (this *Controller) Delete(c *gin.Context) {
	obj := this.GetResModel()
	id, _ := strconv.Atoi(c.Param("id"))
	model.Delete(obj, id)
	helper.ReturnSuccess(c, http.StatusOK, gin.H{"id": id})
}

func (this *Controller) DaoService(funcName string) func(c *gin.Context) {
	if this.GetResModel() == nil {
		panic(model.NotExistDaoError(errors.New("model not exist ")))
	}
	switch funcName {
	case "info":
		return this.Info
	case "create":
		return this.Create
	case "update":
		return this.Update
	case "delete":
		return this.Delete
	default:
		return this.List
	}

}