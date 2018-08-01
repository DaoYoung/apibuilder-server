package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"log"
	"reflect"
	"errors"
)
type ControllerInterface interface {
	CrudService()
}

type Controller struct {
	Res model.Resource

}

func (this *Controller) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info := this.Res.InitDao().ByID(id)
	c.JSON(http.StatusOK, info)
}
func (this *Controller) List(c *gin.Context) {
	log.Println(this.Res.InitDao())
	log.Println(reflect.TypeOf(this.Res.InitDao()))
	list := this.Res.InitDao().FindList()
	c.JSON(http.StatusOK, list)
}
func (this *Controller) Create(c *gin.Context) {
	err := c.BindJSON(this.Res)
	if err != nil {
		panic(JsonTypeError(err))
	}
	info := this.Res.InitDao().Create(this.Res)
	c.JSON(http.StatusCreated, info)
}
func (this *Controller) Update(c *gin.Context) {
	err := c.BindJSON(this.Res)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := this.Res.InitDao().Update(id, this.Res)
	c.JSON(http.StatusOK, info)
}
func (this *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	this.Res.InitDao().Delete(id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (this *Controller) CrudService(funcName string) func(c *gin.Context) {
	if this.Res.InitDao() == nil{
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
